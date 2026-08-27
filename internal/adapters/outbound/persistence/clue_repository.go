package persistence

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/outbound/persistence/sqlc"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Type assertion ClueRepository must implement the ClueRepository port
var _ ports.ClueRepository = &ClueRepository{}

// ClueRepository is a struct that implements the ClueRepository
// port for interacting with the persistence layer
type ClueRepository struct {
	querier *sqlc.Queries
}

// NewClueRepository creates a new instance of ClueRepository with the provided querier
func NewClueRepository(querier *sqlc.Queries) *ClueRepository {
	return &ClueRepository{
		querier: querier,
	}
}

// SaveClue saves a clue entry to the repository related to a specific challenge ID
func (c *ClueRepository) SaveClue(ctx context.Context, challengeId int64, clue *domain.Clue) error {
	// Use the querier to create a new clue in the database
	saved, err := c.querier.CreateClue(ctx, sqlc.CreateClueParams{
		ChallengeID:   challengeId,
		Description:   clue.Description,
		SequenceOrder: int64(clue.SequenceOrder),
	})
	if err != nil {
		return err
	}

	// Reflect the DB-generated ID back onto the domain entity
	clue.Id = saved.ID

	return nil
}

// GetClue retrieves a specific clue associated with a challenge ID from the repository
func (c *ClueRepository) GetClues(ctx context.Context, challengeId int64) ([]*domain.Clue, error) {
	// Use the querier to get clues from the database by challenge ID
	clues, err := c.querier.GetCluesByChallengeId(ctx, challengeId)
	if err != nil {
		return nil, err
	}

	// Map the retrieved clues to the domain model
	var domainClues []*domain.Clue
	for _, clue := range clues {
		domainClue := &domain.Clue{
			Id:            clue.ID,
			Description:   clue.Description,
			SequenceOrder: int(clue.SequenceOrder),
		}
		domainClues = append(domainClues, domainClue)
	}

	return domainClues, nil
}
