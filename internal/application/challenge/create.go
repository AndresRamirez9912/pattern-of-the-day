package challenge

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// CreateChallengeUseCase is responsible for generating a new challenge using
// the LLM provider and saving it to the ChallengeRepository
type CreateChallengeUseCase struct {
	Logger              app.Logger
	LLMProvider         ports.LLMProvider
	ChallengeRepository ports.ChallengeRepository
}

// NewCreateChallengeUseCase creates a new instance of CreateChallengeUseCase with the provided LLMProvider and ChallengeRepository
func NewCreateChallengeUseCase(logger app.Logger, llmProvider ports.LLMProvider, challengeRepository ports.ChallengeRepository) *CreateChallengeUseCase {
	return &CreateChallengeUseCase{
		Logger:              logger,
		LLMProvider:         llmProvider,
		ChallengeRepository: challengeRepository,
	}
}

// Execute generates a new challenge using the LLM provider and saves it to the ChallengeRepository
func (c *CreateChallengeUseCase) Execute(ctx context.Context, req ports.ChallengeGenerationRequest) (*domain.Challenge, error) {
	// Generate a new challenge using the LLM provider
	challenge, err := c.LLMProvider.GenerateChallente(ctx, req)
	if err != nil {
		c.Logger.Error("error generating challenge though the LLM provider", "error", err.Error())
		return nil, err
	}

	// Save the generated challenge to the ChallengeRepository
	err = c.ChallengeRepository.SaveChallenge(challenge)
	if err != nil {
		c.Logger.Error("error saving challenge to the repository", "error", err.Error())
		return nil, err
	}

	return challenge, nil
}
