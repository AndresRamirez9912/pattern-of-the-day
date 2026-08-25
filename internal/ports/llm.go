package ports

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// LLMProvider defines the interface for interacting with a Large Language Model (LLM) provider
type LLMProvider interface {
	// GenerateChallenge generates a new challenge using the LLM provider with the
	GenerateChallente(ctx context.Context, req ChallengeGenerationRequest) (*domain.Challenge, error)
	// EvaluateSolution evaluates a solution for a given challenge using the LLM provider and returns feedback
	EvaluateSolution(ctx context.Context, challenge *domain.Challenge, solutionPath string) (*domain.Feedback, error)
	// GenerateClue generates a clue for a given challenge using the LLM provider and returns the clue
	GenerateClue(ctx context.Context, challenge *domain.Challenge) (*domain.Clue, error)
}

// ChallengeGenerationRequest represents the request to generate a new challenge using the LLM provider
type ChallengeGenerationRequest struct {
	Topic      string
	Difficulty domain.Difficulty
	Pattern    domain.Pattern
	Type       domain.ChallengeType
}
