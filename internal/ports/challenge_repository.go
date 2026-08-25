package ports

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// ChallengeRepository defines the interface for interacting with the DB
// to persist and retrieve Challenge entities
type ChallengeRepository interface {
	// SaveChallenge saves a challenge entity to the repository
	SaveChallenge(ctx context.Context, challenge *domain.Challenge) error
	// GetChallengeByID retrieves a challenge entity from the repository by its ID
	GetChallengeByID(ctx context.Context, id int64) (*domain.Challenge, error)
	// UpdateChallenge updates an existing challenge entity in the repository
	UpdateChallenge(ctx context.Context, challenge *domain.Challenge) error
}
