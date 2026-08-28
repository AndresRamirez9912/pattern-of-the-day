package challenge

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// ChallengeDetails bundles a challenge with its attempts and their feedbacks.
type ChallengeDetails struct {
	Challenge *domain.Challenge
	Attempts  []*domain.Attempt
	Feedbacks []*domain.Feedback
}

// GetChallengeDetailsUseCase retrieves a challenge with its clues and feedbacks.
type GetChallengeDetailsUseCase struct {
	Logger              ports.Logger
	ChallengeRepository ports.ChallengeRepository
	AttemptsRepository  ports.AttemptsRepository
	FeedbackRepository  ports.FeedbackRepository
}

// NewGetChallengeDetailsUseCase creates a new GetChallengeDetailsUseCase.
func NewGetChallengeDetailsUseCase(
	logger ports.Logger,
	challengeRepository ports.ChallengeRepository,
	attemptsRepository ports.AttemptsRepository,
	feedbackRepository ports.FeedbackRepository,
) *GetChallengeDetailsUseCase {
	return &GetChallengeDetailsUseCase{
		Logger:              logger,
		ChallengeRepository: challengeRepository,
		AttemptsRepository:  attemptsRepository,
		FeedbackRepository:  feedbackRepository,
	}
}

// Execute retrieves a challenge (with clues) and every feedback left on its attempts.
func (g *GetChallengeDetailsUseCase) Execute(ctx context.Context, id int64) (*ChallengeDetails, error) {
	found, err := g.ChallengeRepository.GetChallengeByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		g.Logger.Error("challenge not found", "challenge_id", id)
		return nil, domain.NewError("challenge not found", domain.ErrCodeNotFound)
	}
	if err != nil {
		g.Logger.Error("error fetching challenge", "challenge_id", id, "error", err.Error())
		return nil, err
	}

	attempts, err := g.AttemptsRepository.ListAttemptsByChallengeId(ctx, id)
	if err != nil {
		g.Logger.Error("error listing attempts", "challenge_id", id, "error", err.Error())
		return nil, err
	}

	// Only completed/failed attempts have feedback attached.
	feedbacks := make([]*domain.Feedback, 0, len(attempts))
	for _, a := range attempts {
		if a.FeedbackId == nil {
			continue
		}

		fb, err := g.FeedbackRepository.GetFeedback(ctx, *a.FeedbackId)
		if err != nil {
			g.Logger.Error("error fetching feedback", "feedback_id", *a.FeedbackId, "error", err.Error())
			return nil, err
		}

		feedbacks = append(feedbacks, fb)
	}

	return &ChallengeDetails{Challenge: found, Attempts: attempts, Feedbacks: feedbacks}, nil
}
