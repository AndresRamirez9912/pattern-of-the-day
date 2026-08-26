package persistence

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/persistence/sqlc"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Type assertion
var _ ports.ChallengeRepository = &ChallengeRepository{}

// ChallengeRepository is a struct that implements the ChallengeRepository port
// for interacting with the persistence layer.
type ChallengeRepository struct {
	querier *sqlc.Queries
}

// NewChallengeRepository creates a new instance of ChallengeRepository with the provided querier
func NewChallengeRepository(querier *sqlc.Queries) *ChallengeRepository {
	return &ChallengeRepository{
		querier: querier,
	}
}

// SaveChallenge saves a challenge entity to the repository
func (c *ChallengeRepository) SaveChallenge(ctx context.Context, challenge *domain.Challenge) error {
	// Use the querier to create a new challenge in the database
	_, err := c.querier.CreateChallenge(ctx, sqlc.CreateChallengeParams{
		Name:          challenge.Name,
		Description:   challenge.Description,
		Difficulty:    int64(challenge.Dificulty),
		Type:          string(challenge.Type),
		TargetPattern: string(challenge.Pattern),
		UserID:        challenge.UserId,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetChallengeByID retrieves a challenge entity from the repository by its ID
func (c *ChallengeRepository) GetChallengeByID(ctx context.Context, id int64) (*domain.Challenge, error) {
	// Use the querier to get a challenge from the database by its ID
	challenge, err := c.querier.GetChallengeById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Map the retrieved challenge to the domain model
	return &domain.Challenge{
		Id:          challenge.ID,
		Name:        challenge.Name,
		Description: challenge.Description,
		Dificulty:   domain.Difficulty(challenge.Difficulty),
		Type:        domain.ChallengeType(challenge.Type),
	}, nil
}

// UpdateChallenge updates an existing challenge entity in the repository
func (c *ChallengeRepository) UpdateChallenge(ctx context.Context, challenge *domain.Challenge) error {
	// Use the querier to update the challenge in the database
	_, err := c.querier.UpdateChallenge(ctx, sqlc.UpdateChallengeParams{
		ID:          challenge.Id,
		Name:        challenge.Name,
		Description: challenge.Description,
		Difficulty:  int64(challenge.Dificulty),
		Type:        string(challenge.Type),
	})
	if err != nil {
		return err
	}

	return nil
}
