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
}

// NewProvider creates a new Ollama-backed LLMProvider.
func NewProvider(client *Client, model string) *Provider {
	return &Provider{
		Client: client,
		Model:  model,
	}
}

// GenerateChallente generates a new challenge using the LLM provider
func (o *Provider) GenerateChallente(ctx context.Context, req ports.ChallengeGenerationRequest) (*domain.Challenge, error) {
	resp, err := o.Client.Chat(ctx, ChatRequest{
		Model:  o.Model,
		Format: "json",
		Messages: []ChatMessage{
			{Role: "system", Content: challengeSystemPrompt},
			{Role: "user", Content: challengeUserPrompt(req)},
		},
	})
	if err != nil {
		return nil, err
	}

	parsed, err := decodeJSON[challengeResponse](resp.Message.Content)
	if err != nil {
		return nil, err
	}

	return domain.NewChallenge(0, parsed.Name, parsed.Description, req.Difficulty, req.Type, req.Pattern, req.UserId), nil
}

// EvaluateSolution evaluates a solution for a given challenge using the LLM provider and returns feedback.
// solutionPath may point to a single source file or to a directory containing a whole project.
func (o *Provider) EvaluateSolution(ctx context.Context, attempt *domain.Attempt, challenge *domain.Challenge, solutionPath string) (*domain.Feedback, error) {
	solutionCode, err := readSolutionSource(solutionPath)
	if err != nil {
		return nil, err
	}

	resp, err := o.Client.Chat(ctx, ChatRequest{
		Model:  o.Model,
		Format: "json",
		Messages: []ChatMessage{
			{Role: "system", Content: feedbackSystemPrompt},
			{Role: "user", Content: feedbackUserPrompt(challenge, solutionCode)},
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
	resp, err := o.Client.Chat(ctx, ChatRequest{
		Model:  o.Model,
		Format: "json",
		Messages: []ChatMessage{
			{Role: "system", Content: clueSystemPrompt},
			{Role: "user", Content: clueUserPrompt(challenge)},
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
