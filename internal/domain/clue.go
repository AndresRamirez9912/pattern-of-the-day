package domain

// Clue represents a clue for a challenge
type Clue struct {
	Description string
}

// NewClue creates a new clue with the given description
func NewClue(description string) *Clue {
	return &Clue{
		Description: description,
	}
}
