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
	attemptRepository  ports.AttemptsRepository
	fileWriter         ports.FileWriter
}

// NewCreateFeedbackUseCase creates a new instance of CreateFeedbackUseCase
func NewCreateFeedbackUseCase(
	logger ports.Logger,
	feedbackRepository ports.FeedbackRepository,
	llmProvider ports.LLMProvider,
	attemptRepository ports.AttemptsRepository,
	fileWriter ports.FileWriter,
) *CreateFeedbackUseCase {
	return &CreateFeedbackUseCase{
		logger:             logger,
		feedbackRepository: feedbackRepository,
		llmProvider:        llmProvider,
		attemptRepository:  attemptRepository,
		fileWriter:         fileWriter,
	}
}

// Execute generates feedback for a given attempt and solution path, saves it to the repository,
// links it to the attempt, marks the attempt as completed, and writes attempt-<N>-feedback.md to outDir.
func (c *CreateFeedbackUseCase) Execute(ctx context.Context, attempt *domain.Attempt, challenge *domain.Challenge, solutionPath, outDir string) (*domain.Feedback, error) {
	// Generate feedback using the LLM provider
	feedback, err := c.llmProvider.EvaluateSolution(ctx, attempt, challenge, solutionPath)
	if err != nil {
		c.logger.Error("failed to generate feedback", "error", err.Error())
		return nil, err
	}

	// Save the feedback to the repository
	err = c.feedbackRepository.SaveFeedback(ctx, feedback)
	if err != nil {
		c.logger.Error("failed to create feedback", "error", err.Error())
		return nil, err
	}

	// Link the feedback to the attempt and mark it as completed
	attempt.FeedbackId = &feedback.Id
	err = attempt.Complete()
	if err != nil {
		c.logger.Error("failed to complete attempt", "attempt_id", attempt.Id, "error", err.Error())
		return nil, err
	}

	// Update the attempt in the repository after linking the feedback and marking it as completed
	err = c.attemptRepository.UpdateAttempt(ctx, attempt)
	if err != nil {
		c.logger.Error("failed to update attempt after feedback", "attempt_id", attempt.Id, "error", err.Error())
		return nil, err
	}

	// Write the feedback to a file in the specified output directory
	err = c.fileWriter.WriteFeedback(ctx, outDir, attempt, challenge, feedback)
	if err != nil {
		c.logger.Error("error writing feedback file", "error", err.Error())
		return nil, err
	}

	return feedback, nil
}
