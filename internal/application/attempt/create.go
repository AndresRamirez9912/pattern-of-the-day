package attempt

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/utils"
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

// Execute creates a new attempt with the given parameters and saves it to the repository
func (c *CreateAttemptUseCase) Execute(ctx context.Context, challengeId int) (*domain.Attempt, error) {
	// Create a new attempt entity as pending and without a feedback refernece
	attempt := &domain.Attempt{
		ChallengeId: utils.ToPtr(int64(challengeId)),
		FeedbackId:  nil,
		Status:      domain.AttemptStatusPending,
	}

	// Save the attempt to the repository
	err := c.AttemptRepository.CreateAttempt(ctx, attempt)
	if err != nil {
		c.Logger.Error("Failed to create attempt", "error", err)
		return nil, err
	}

	return attempt, nil
}
