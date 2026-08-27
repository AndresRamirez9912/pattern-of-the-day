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
	saved, err := c.querier.CreateChallenge(ctx, sqlc.CreateChallengeParams{
		Name:        challenge.Name,
		Description: challenge.Description,
		Difficulty:  string(challenge.Difficulty),
		Type:        string(challenge.Type),
		Target:      challenge.Target,
		UserID:      challenge.UserId,
	})
	if err != nil {
		return err
	}

	// Reflect the DB-generated ID back onto the domain entity
	challenge.Id = saved.ID

	return nil
}

// GetChallengeByID retrieves a challenge entity from the repository by its ID
func (c *ChallengeRepository) GetChallengeByID(ctx context.Context, id int64) (*domain.Challenge, error) {
	// Use the querier to get a challenge from the database by its ID
	challenge, err := c.querier.GetChallengeById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Fetch the clues associated with the challenge
	clues, err := c.querier.GetCluesByChallengeId(ctx, id)
	if err != nil {
		return nil, err
	}

	domainClues := make([]domain.Clue, 0, len(clues))
	for _, clue := range clues {
		domainClues = append(domainClues, domain.Clue{
			Id:            clue.ID,
			Description:   clue.Description,
			SequenceOrder: int(clue.SequenceOrder),
		})
	}

	// Map the retrieved challenge to the domain model
	return &domain.Challenge{
		Id:          challenge.ID,
		Name:        challenge.Name,
		Description: challenge.Description,
		Difficulty:  domain.Difficulty(challenge.Difficulty),
		Type:        domain.ChallengeType(challenge.Type),
		Target:      challenge.Target,
		UserId:      challenge.UserID,
		Clues:       domainClues,
	}, nil
}

// UpdateChallenge updates an existing challenge entity in the repository
func (c *ChallengeRepository) UpdateChallenge(ctx context.Context, challenge *domain.Challenge) error {
	// Use the querier to update the challenge in the database
	_, err := c.querier.UpdateChallenge(ctx, sqlc.UpdateChallengeParams{
		ID:          challenge.Id,
		Name:        challenge.Name,
		Description: challenge.Description,
		Difficulty:  string(challenge.Difficulty),
		Type:        string(challenge.Type),
		Target:      challenge.Target,
	})
	if err != nil {
		return err
	}

	return nil
}
