package domain

// Clue represents a clue for a challenge
type Clue struct {
	Id            int64
	Description   string
	SequenceOrder int
}

// NewClue creates a new clue with the given description
func NewClue(id int64, description string, sequence int) *Clue {
	return &Clue{
		Id:            id,
		Description:   description,
		SequenceOrder: sequence,
	}
}
