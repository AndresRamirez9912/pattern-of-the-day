package ollama

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Type assertion Ollama must implement the LLMProvider port
var _ ports.LLMProvider = &Provider{}

// Provider is a concrete implementation of the LLMProvider interface backed by Ollama.
type Provider struct {
	Client *Client
	Model  string
	Logger ports.Logger
}

// NewProvider creates a new Ollama-backed LLMProvider.
func NewProvider(client *Client, model string, logger ports.Logger) *Provider {
	return &Provider{
		Client: client,
		Model:  model,
		Logger: logger,
	}
}

// GenerateChallenge generates a new challenge using the LLM provider
func (o *Provider) GenerateChallenge(ctx context.Context, req ports.ChallengeGenerationRequest) (*domain.Challenge, error) {
	o.Logger.Info("generating challenge via LLM provider", "model", o.Model, "type", string(req.Type), "target", req.Target, "difficulty", string(req.Difficulty))

	system := challengeSystemPrompt
	user := challengeUserPrompt(req)

	resp, err := o.Client.Chat(ctx, ChatRequest{
		Model:   o.Model,
		Format:  "json",
		Options: contextOptionsFor(system, user),
		Messages: []ChatMessage{
			{Role: "system", Content: system},
			{Role: "user", Content: user},
		},
	})
	if err != nil {
		return nil, err
	}

	parsed, err := decodeJSON[challengeResponse](resp.Message.Content)
	if err != nil {
		return nil, err
	}

	return domain.NewChallenge(0, parsed.Name, parsed.Description, req.Difficulty, req.Type, req.Target, req.UserId), nil
}

// EvaluateSolution evaluates a solution for a given challenge using the LLM provider and returns feedback.
// solutionPath may point to a single source file or to a directory containing a whole project.
func (o *Provider) EvaluateSolution(ctx context.Context, attempt *domain.Attempt, challenge *domain.Challenge, solutionPath string) (*domain.Feedback, error) {
	solutionCode, err := readSolutionSource(solutionPath)
	if err != nil {
		return nil, err
	}

	o.Logger.Info("evaluating solution via LLM provider", "model", o.Model, "challenge_id", challenge.Id, "attempt_id", attempt.Id, "solution_chars", len(solutionCode))

	system := feedbackSystemPrompt
	user := feedbackUserPrompt(challenge, solutionCode)

	resp, err := o.Client.Chat(ctx, ChatRequest{
		Model:   o.Model,
		Format:  "json",
		Options: contextOptionsFor(system, user),
		Messages: []ChatMessage{
			{Role: "system", Content: system},
			{Role: "user", Content: user},
		},
	})
	if err != nil {
		return nil, err
	}

	parsed, err := decodeJSON[feedbackResponse](resp.Message.Content)
	if err != nil {
		return nil, err
	}

	feedback := domain.NewFeedback(parsed.Score, parsed.Summary)
	for _, suggestion := range parsed.Suggestions {
		feedback.AddSuggestion(suggestion)
	}

	return feedback, nil
}

// GenerateClue generates a clue for a given challenge using the LLM provider and returns the clue
func (o *Provider) GenerateClue(ctx context.Context, challenge *domain.Challenge) (*domain.Clue, error) {
	o.Logger.Info("generating clue via LLM provider", "model", o.Model, "challenge_id", challenge.Id, "clue_number", len(challenge.Clues)+1)

	system := clueSystemPrompt
	user := clueUserPrompt(challenge)

	resp, err := o.Client.Chat(ctx, ChatRequest{
		Model:   o.Model,
		Format:  "json",
		Options: contextOptionsFor(system, user),
		Messages: []ChatMessage{
			{Role: "system", Content: system},
			{Role: "user", Content: user},
		},
	})
	if err != nil {
		return nil, err
	}

	parsed, err := decodeJSON[clueResponse](resp.Message.Content)
	if err != nil {
		return nil, err
	}

	return domain.NewClue(0, parsed.Clue, len(challenge.Clues)+1), nil
}
