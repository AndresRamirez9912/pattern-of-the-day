package ports

// ClueRepository defines the interface for storing and retrieving clues
type ClueRepository interface {
	// SaveClue saves a clue entry to the repository related to a specific challenge ID
	SaveClue(challengeId, clue string) error
	// GetClue retrieves a specific clue associated with a challenge ID from the repository
	GetClues(challengeId string) ([]string, error)
}
