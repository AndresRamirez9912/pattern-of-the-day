package challenge

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// CreateChallengeUseCase is responsible for generating a new challenge using
// the LLM provider and saving it to the ChallengeRepository
type CreateChallengeUseCase struct {
	Logger              ports.Logger
	LLMProvider         ports.LLMProvider
	ChallengeRepository ports.ChallengeRepository
	AttemptsRepository  ports.AttemptsRepository
	UserRepository      ports.UserRepository
}

// NewCreateChallengeUseCase creates a new instance of CreateChallengeUseCase with the provided LLMProvider and ChallengeRepository
func NewCreateChallengeUseCase(
	logger ports.Logger,
	llmProvider ports.LLMProvider,
	challengeRepository ports.ChallengeRepository,
	attemptsRepository ports.AttemptsRepository,
	userRepository ports.UserRepository,
) *CreateChallengeUseCase {
	return &CreateChallengeUseCase{
		Logger:              logger,
		LLMProvider:         llmProvider,
		ChallengeRepository: challengeRepository,
		AttemptsRepository:  attemptsRepository,
		UserRepository:      userRepository,
	}
}

// Execute generates a new challenge using the LLM provider and saves it to the ChallengeRepository
func (c *CreateChallengeUseCase) Execute(ctx context.Context, req ports.ChallengeGenerationRequest) (*domain.Challenge, error) {
	// Validate the user received (must exists)
	user, err := c.UserRepository.GetUserByID(ctx, req.UserId)
	if errors.Is(err, sql.ErrNoRows) {
		c.Logger.Error("user not found", "user_id", req.UserId)
		return nil, errors.New("user not found")
	}
	if err != nil {
		c.Logger.Error("error fetching user", "error", err.Error())
		return nil, err
	}

	// Generate a new challenge using the LLM provider
	challenge, err := c.LLMProvider.GenerateChallente(ctx, req)
	c.Logger.Info("Generating challenge...")
	if err != nil {
		c.Logger.Error("error generating challenge though the LLM provider", "error", err.Error())
		return nil, err
	}

	// Assign the user Id to the generated challenge
	challenge.UserId = user.Id

	// Save the generated challenge to the ChallengeRepository
	err = c.ChallengeRepository.SaveChallenge(ctx, challenge)
	if err != nil {
		c.Logger.Error("error saving challenge to the repository", "error", err.Error())
		return nil, err
	}

	// Create a new attempt for the generated challenge
	err = c.AttemptsRepository.CreateAttempt(ctx, &domain.Attempt{
		ChallengeId: &challenge.Id,
		Status:      domain.AttemptStatusPending,
	})
	if err != nil {
		c.Logger.Error("error creating attempt for the challenge", "error", err.Error())
		return nil, err
	}

	return challenge, nil
}
