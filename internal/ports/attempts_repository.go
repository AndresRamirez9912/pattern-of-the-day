package ports

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// AttemptsRepository defines the interface for attempt repository operations
type AttemptsRepository interface {
	// CreateAttempt creates a new attempt entity in the repository
	CreateAttempt(ctx context.Context, attempt *domain.Attempt) error
	// UpdateAttempt updates an existing attempt entity in the repository
	UpdateAttempt(ctx context.Context, attempt *domain.Attempt) error
	// GetAttemptByID retrieves an attempt entity from the repository by its ID
	GetAttemptByID(ctx context.Context, id int64) (*domain.Attempt, error)
	// ListAttemptsByChallengeId retrieves every attempt entity associated with a challenge
	ListAttemptsByChallengeId(ctx context.Context, challengeId int64) ([]*domain.Attempt, error)
	// ListAttemptsByUserChallenge retrieves all attempt entities for a challenge owned by a specific user
	ListAttemptsByUserChallenge(ctx context.Context, userId, challengeId int64) ([]*domain.Attempt, error)
}
