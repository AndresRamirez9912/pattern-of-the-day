package ports

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// ClueRepository defines the interface for storing and retrieving clues
type ClueRepository interface {
	// SaveClue saves a clue entry to the repository related to a specific challenge ID
	SaveClue(ctx context.Context, challengeId int64, clue *domain.Clue) error
	// GetClue retrieves a specific clue associated with a challenge ID from the repository
	GetClues(ctx context.Context, challengeId int64) ([]*domain.Clue, error)
}
