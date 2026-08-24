package domain

// Feedback represents the feedback given to a user after completing a challenge or task
// It contains information about the score, summary, and suggestions for improvement.
type Feedback struct {
	// Score is the score given to the feedback
	Score int
	// Summary is a brief summary of the feedback
	Summary string
	// Suggestions is a list of suggestions for improvement
	Suggestions []string
}

// NewFeedback creates a new Feedback instance with the given score and summary
func NewFeedback(score int, summary string) *Feedback {
	return &Feedback{
		Score:       score,
		Summary:     summary,
		Suggestions: []string{},
	}
}

// AddSuggestion adds a suggestion for improvement to the feedback
func (f *Feedback) AddSuggestion(suggestion string) {
	f.Suggestions = append(f.Suggestions, suggestion)
}

// HasSuggestions checks if the feedback has any suggestions for improvement
func (f *Feedback) HasSuggestions() bool {
	return len(f.Suggestions) > 0
}
