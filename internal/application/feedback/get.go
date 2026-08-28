package feedback

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// GetFeedbackUseCase retrieves feedback by ID.
type GetFeedbackUseCase struct {
	Logger             ports.Logger
	FeedbackRepository ports.FeedbackRepository
}

// NewGetFeedbackUseCase creates a new GetFeedbackUseCase.
func NewGetFeedbackUseCase(logger ports.Logger, feedbackRepository ports.FeedbackRepository) *GetFeedbackUseCase {
	return &GetFeedbackUseCase{
		Logger:             logger,
		FeedbackRepository: feedbackRepository,
	}
}

// Execute retrieves feedback by its ID.
func (g *GetFeedbackUseCase) Execute(ctx context.Context, id int64) (*domain.Feedback, error) {
	found, err := g.FeedbackRepository.GetFeedback(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		g.Logger.Error("feedback not found", "feedback_id", id)
		return nil, domain.NewError("feedback not found", domain.ErrCodeNotFound)
	}
	if err != nil {
		g.Logger.Error("error fetching feedback", "feedback_id", id, "error", err.Error())
		return nil, err
	}

	return found, nil
}
