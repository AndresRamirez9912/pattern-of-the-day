package challenge

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
	FileWriter          ports.FileWriter
}

// NewCreateChallengeUseCase creates a new instance of CreateChallengeUseCase with the provided dependencies
func NewCreateChallengeUseCase(
	logger ports.Logger,
	llmProvider ports.LLMProvider,
	challengeRepository ports.ChallengeRepository,
	attemptsRepository ports.AttemptsRepository,
	userRepository ports.UserRepository,
	fileWriter ports.FileWriter,
) *CreateChallengeUseCase {
	return &CreateChallengeUseCase{
		Logger:              logger,
		LLMProvider:         llmProvider,
		ChallengeRepository: challengeRepository,
		AttemptsRepository:  attemptsRepository,
		UserRepository:      userRepository,
		FileWriter:          fileWriter,
	}
}

// Execute generates a new challenge using the LLM provider, saves it to the ChallengeRepository,
// creates the initial pending attempt for it, and writes challenge.md to outDir.
func (c *CreateChallengeUseCase) Execute(
	ctx context.Context,
	userName string,
	req ports.ChallengeGenerationRequest,
	outDir string,
) (*domain.Challenge, *domain.Attempt, error) {
	// Validate the challenge type received
	if !domain.IsValidChallengeType(req.Type) {
		return nil, nil, fmt.Errorf("invalid challenge type %q", req.Type)
	}

	if req.Target == "" {
		req.Target = randomTarget(req.Type)
	}

	if req.Topic == "" {
		req.Topic = randomTopic()
	}

	if !domain.IsValidDifficulty(req.Difficulty) {
		return nil, nil, fmt.Errorf("invalid difficulty %q", req.Difficulty)
	}

	c.Logger.Info("generating challenge", "topic", req.Topic, "type", req.Type, "difficulty", req.Difficulty)

	// Validate the user received (must exist)
	user, err := c.UserRepository.GetUserByUsername(ctx, userName)
	if errors.Is(err, sql.ErrNoRows) {
		c.Logger.Error("user not found", "username", userName)
		return nil, nil, errors.New("user not found")
	}
	if err != nil {
		c.Logger.Error("error fetching user", "error", err.Error())
		return nil, nil, err
	}

	// Generate a new challenge using the LLM provider
	challenge, err := c.LLMProvider.GenerateChallenge(ctx, req)
	if err != nil {
		c.Logger.Error("error generating challenge though the LLM provider", "error", err.Error())
		return nil, nil, err
	}

	// Assign the user Id to the generated challenge
	challenge.UserId = user.Id

	// Save the generated challenge to the ChallengeRepository
	err = c.ChallengeRepository.SaveChallenge(ctx, challenge)
	if err != nil {
		c.Logger.Error("error saving challenge to the repository", "error", err.Error())
		return nil, nil, err
	}

	// Create the initial pending attempt for the generated challenge
	attempt := &domain.Attempt{
		ChallengeId:   &challenge.Id,
		Status:        domain.AttemptStatusPending,
		SequenceOrder: 1,
	}
	err = c.AttemptsRepository.CreateAttempt(ctx, attempt)
	if err != nil {
		c.Logger.Error("error creating attempt for the challenge", "error", err.Error())
		return nil, nil, err
	}

	// Write the challenge details to challenge.md
	err = c.FileWriter.WriteChallenge(ctx, outDir, challenge)
	if err != nil {
		c.Logger.Error("error writing challenge file", "error", err.Error())
		return nil, nil, err
	}

	return challenge, attempt, nil
}
