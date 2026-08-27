package persistence

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/persistence/sqlc"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Type assertion AttemptsRepository must implement the AttemptsRepository port
var _ ports.AttemptsRepository = &AttemptsRepository{}

// AttemptsRepository is a concrete implementation of the AttemptsRepository interface
type AttemptsRepository struct {
	querier *sqlc.Queries
}

// NewAttemptsRepository creates a new instance of AttemptsRepository with the provided querier
func NewAttemptsRepository(querier *sqlc.Queries) *AttemptsRepository {
	return &AttemptsRepository{
		querier: querier,
	}
}

// CreateAttempt creates a new attempt entity in the repository
func (a *AttemptsRepository) CreateAttempt(ctx context.Context, attempt *domain.Attempt) error {
	// Use the querier to create a new attempt in the database
	saved, err := a.querier.CreateAttempt(ctx, sqlc.CreateAttemptParams{
		ChallengeID:   *attempt.ChallengeId,
		FeedbackID:    attempt.FeedbackId,
		Status:        string(attempt.Status),
		SequenceOrder: int64(attempt.SequenceOrder),
	})
	if err != nil {
		return err
	}

	// Reflect the DB-generated ID back onto the domain entity
	attempt.Id = saved.ID

	return nil
}

// UpdateAttempt updates an existing attempt entity in the repository
func (a *AttemptsRepository) UpdateAttempt(ctx context.Context, attempt *domain.Attempt) error {
	// Use the querier to update the attempt in the database
	_, err := a.querier.UpdateAttempt(ctx, sqlc.UpdateAttemptParams{
		ID:         attempt.Id,
		FeedbackID: attempt.FeedbackId,
		Status:     string(attempt.Status),
	})
	if err != nil {
		return err
	}

	return nil
}

// GetAttemptByID retrieves an attempt entity from the repository by its ID
func (a *AttemptsRepository) GetAttemptByID(ctx context.Context, id int64) (*domain.Attempt, error) {
	// Use the querier to get an attempt from the database by its ID
	attempt, err := a.querier.GetAttemptById(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapAttempt(attempt), nil
}

// ListAttemptsByChallengeId retrieves every attempt entity associated with a challenge
func (a *AttemptsRepository) ListAttemptsByChallengeId(ctx context.Context, challengeId int64) ([]*domain.Attempt, error) {
	// Use the querier to get attempts from the database by challenge ID
	attempts, err := a.querier.ListAttemptsByChallengeId(ctx, challengeId)
	if err != nil {
		return nil, err
	}

	domainAttempts := make([]*domain.Attempt, 0, len(attempts))
	for _, attempt := range attempts {
		domainAttempts = append(domainAttempts, mapAttempt(attempt))
	}

	return domainAttempts, nil
}

// ListAttemptsByUserChallenge retrieves all attempt entities for a challenge owned by a specific user
func (a *AttemptsRepository) ListAttemptsByUserChallenge(ctx context.Context, userId, challengeId int64) ([]*domain.Attempt, error) {
	// Use the querier to get attempts from the database by user and challenge ID
	attempts, err := a.querier.ListAttemptsByUserChallenge(ctx, sqlc.ListAttemptsByUserChallengeParams{
		UserID: userId,
		ID:     challengeId,
	})
	if err != nil {
		return nil, err
	}

	domainAttempts := make([]*domain.Attempt, 0, len(attempts))
	for _, row := range attempts {
		domainAttempts = append(domainAttempts, &domain.Attempt{
			Id:            row.ID,
			ChallengeId:   &row.ChallengeID,
			FeedbackId:    row.FeedbackID,
			Status:        domain.AttemptStatus(row.Status),
			SequenceOrder: int(row.SequenceOrder),
		})
	}

	return domainAttempts, nil
}

// mapAttempt converts a generated sqlc.Attempt row into a domain.Attempt.
func mapAttempt(attempt sqlc.Attempt) *domain.Attempt {
	return &domain.Attempt{
		Id:            attempt.ID,
		ChallengeId:   &attempt.ChallengeID,
		FeedbackId:    attempt.FeedbackID,
		Status:        domain.AttemptStatus(attempt.Status),
		SequenceOrder: int(attempt.SequenceOrder),
	}
}
