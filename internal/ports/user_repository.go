package ports

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	// SaveUser saves a user entity to the repository
	SaveUser(ctx context.Context, user *domain.User) error
	// GetUserByID retrieves a user entity from the repository by its ID
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	// ListUserChallenges retrieves all challenges associated with a user
	ListUserChallenges(ctx context.Context, userId int64) ([]*domain.Challenge, error)
}
