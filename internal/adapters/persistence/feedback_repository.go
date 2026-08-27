package persistence

import (
	"context"
	"encoding/json"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/persistence/sqlc"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Type assertion FeedbackRepository implements FeedbackRepository from ports
var _ ports.FeedbackRepository = &FeedbackRepository{}

// FeedbackRepository is the implementation of the FeedbackRepository interface from the ports package.
type FeedbackRepository struct {
	querier *sqlc.Queries
}

// NewFeedbackRepository creates a new instance of FeedbackRepository with the given querier.
func NewFeedbackRepository(querier *sqlc.Queries) *FeedbackRepository {
	return &FeedbackRepository{
		querier: querier,
	}
}

// SaveFeedback saves a feedback entry to the repository
func (f *FeedbackRepository) SaveFeedback(ctx context.Context, feedback *domain.Feedback) error {
	// Marshall suggestions to JSON
	suggestionsJSON, err := json.Marshal(feedback.Suggestions)
	if err != nil {
		return err
	}

	// Use the querier to insert the feedback into the database
	saved, err := f.querier.CreateFeedback(ctx, sqlc.CreateFeedbackParams{
		Suggestions: string(suggestionsJSON),
		Summary:     feedback.Summary,
		Score:       int64(feedback.Score),
	})
	if err != nil {
		return err
	}

	// Reflect the DB-generated ID back onto the domain entity
	feedback.Id = saved.ID

	return nil
}

// GetFeedback retrieves a feedback entry from the repository
func (f *FeedbackRepository) GetFeedback(ctx context.Context, feedbackId int64) (*domain.Feedback, error) {
	// Use the querier to retrieve the feedback entry from the database by its ID
	feedback, err := f.querier.GetFeedbackById(ctx, feedbackId)
	if err != nil {
		return nil, err
	}

	// Unmarshal suggestions from JSON
	var suggestions []string
	err = json.Unmarshal([]byte(feedback.Suggestions), &suggestions)
	if err != nil {
		return nil, err
	}

	// Return the feedback entry as a domain feedback entity
	return &domain.Feedback{
		Id:          feedback.ID,
		Suggestions: suggestions,
		Score:       int(feedback.Score),
		Summary:     feedback.Summary,
	}, nil
}
