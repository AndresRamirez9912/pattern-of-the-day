package attempt

import (
	"context"
	"fmt"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// CreateAttemptUseCase handles the creation of a new attempt in the application
type CreateAttemptUseCase struct {
	Logger            ports.Logger
	AttemptRepository ports.AttemptsRepository
}

// NewCreateAttemptUseCase creates a new instance of CreateAttemptUseCase
func NewCreateAttemptUseCase(logger ports.Logger, attemptRepository ports.AttemptsRepository) *CreateAttemptUseCase {
	return &CreateAttemptUseCase{
		Logger:            logger,
		AttemptRepository: attemptRepository,
	}
}

// Execute creates a new attempt for the given challenge, as long as every
// attempt already associated with that challenge is closed (completed or
// failed) — a challenge can only have one attempt in progress at a time.
func (c *CreateAttemptUseCase) Execute(ctx context.Context, challengeId int64) (*domain.Attempt, error) {
	existingAttempts, err := c.AttemptRepository.ListAttemptsByChallengeId(ctx, challengeId)
	if err != nil {
		c.Logger.Error("error listing existing attempts for challenge", "challenge_id", challengeId, "error", err.Error())
		return nil, err
	}

	// Ensure there are no open attempts for the challenge before creating a new one
	for _, existing := range existingAttempts {
		if !existing.IsClosed() {
			return nil, fmt.Errorf("challenge %d already has an open attempt (id %d); finish it before starting a new one", challengeId, existing.Id)
		}
	}

	// Create a new attempt entity as pending and without a feedback reference
	attempt := &domain.Attempt{
		ChallengeId:   &challengeId,
		Status:        domain.AttemptStatusPending,
		SequenceOrder: len(existingAttempts) + 1,
	}

	// Save the attempt to the repository
	err = c.AttemptRepository.CreateAttempt(ctx, attempt)
	if err != nil {
		c.Logger.Error("failed to create attempt", "error", err.Error())
		return nil, err
	}

	return attempt, nil
}
