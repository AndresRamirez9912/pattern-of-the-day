package ports

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// AttemptsRepository defines the interface for attempt repository operations
type AttemptsRepository interface {
	// CreateAttempt creates a new attempt entity in the repository
	CreateAttempt(ctx context.Context, attempt *domain.Attempt) error
	// GetAttemptByID retrieves an attempt entity from the repository by its ID
	GetAttemptByID(ctx context.Context, id int) (*domain.Attempt, error)
	// ListAttemptsByChallenge retrieves all attempt entities from the repository
	ListAttemptsByUserChallenge(ctx context.Context, userId, challengeId int) ([]*domain.Attempt, error)
}
