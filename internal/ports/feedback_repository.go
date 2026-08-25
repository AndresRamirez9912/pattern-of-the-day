package ports

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
)

// FeedbackRepository defines the interface for storing and retrieving feedback entries
type FeedbackRepository interface {
	// SaveFeedback saves a feedback entry to the repository
	SaveFeedback(ctx context.Context, feedback *domain.Feedback) error
	// GetFeedbacks retrieves all feedback entries from the repository
	GetFeedbacks(ctx context.Context) ([]*domain.Feedback, error)
}
