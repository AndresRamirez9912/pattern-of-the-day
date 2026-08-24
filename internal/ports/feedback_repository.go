package ports

// FeedbackRepository defines the interface for storing and retrieving feedback entries
type FeedbackRepository interface {
	// SaveFeedback saves a feedback entry to the repository
	SaveFeedback(challengeId, feedback string) error
	// GetChallengeFeedback retrieves all feedback entries associated with a specific challenge ID from the repository
	GetChallengeFeedback(challengeID string) ([]string, error)
	// GetFeedbacks retrieves all feedback entries from the repository/
	GetFeedbacks() ([]string, error)
}
