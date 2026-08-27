package feedback

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// CreateFeedbackUseCase handles the creation of a new feedback in the application
type CreateFeedbackUseCase struct {
	logger             ports.Logger
	feedbackRepository ports.FeedbackRepository
	llmProvider        ports.LLMProvider
}

// NewCreateFeedbackUseCase creates a new instance of CreateFeedbackUseCase
func NewCreateFeedbackUseCase(logger ports.Logger, feedbackRepository ports.FeedbackRepository, llmProvider ports.LLMProvider) *CreateFeedbackUseCase {
	return &CreateFeedbackUseCase{
		logger:             logger,
		feedbackRepository: feedbackRepository,
		llmProvider:        llmProvider,
	}
}

// Execute generates feedback for a given attempt and solution path, saves it to the repository, and returns the feedback.
func (c *CreateFeedbackUseCase) Execute(ctx context.Context, attempt *domain.Attempt, challenge *domain.Challenge, solutionPath string) (*domain.Feedback, error) {
	// Generate feedback using the LLM provider
	feedback, err := c.llmProvider.EvaluateSolution(ctx, attempt, challenge, solutionPath)
	if err != nil {
		c.logger.Error("Failed to generate feedback", "error", err)
		return nil, err
	}

	// Save the feedback to the repository
	err = c.feedbackRepository.SaveFeedback(ctx, feedback)
	if err != nil {
		c.logger.Error("Failed to create feedback", "error", err)
		return nil, err
	}

	return feedback, nil
}
