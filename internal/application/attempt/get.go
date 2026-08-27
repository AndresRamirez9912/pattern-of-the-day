package attempt

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// GetAttemptUseCase retrieves a single attempt by ID.
type GetAttemptUseCase struct {
	Logger            ports.Logger
	AttemptRepository ports.AttemptsRepository
}

// NewGetAttemptUseCase creates a new instance of GetAttemptUseCase
func NewGetAttemptUseCase(logger ports.Logger, attemptRepository ports.AttemptsRepository) *GetAttemptUseCase {
	return &GetAttemptUseCase{
		Logger:            logger,
		AttemptRepository: attemptRepository,
	}
}

// Execute retrieves an attempt by its ID from the AttemptRepository
func (g *GetAttemptUseCase) Execute(ctx context.Context, id int64) (*domain.Attempt, error) {
	attempt, err := g.AttemptRepository.GetAttemptByID(ctx, id)
	if err != nil {
		g.Logger.Error("error fetching attempt", "attempt_id", id, "error", err.Error())
		return nil, err
	}

	return attempt, nil
}
