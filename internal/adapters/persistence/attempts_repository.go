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
	_, err := a.querier.CreateAttempt(ctx, sqlc.CreateAttemptParams{
		ChallengeID: *attempt.ChallengeId,
		FeedbackID:  attempt.FeedbackId,
		Status:      string(attempt.Status),
	})
	if err != nil {
		return err
	}

	return nil
}

// GetAttemptByID retrieves an attempt entity from the repository by its ID
func (a *AttemptsRepository) GetAttemptByID(ctx context.Context, id int) (*domain.Attempt, error) {
	// Use the querier to get an attempt from the database by its ID
	attempt, err := a.querier.GetAttemptById(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	// Map the retrieved attempt to the domain model
	return &domain.Attempt{
		Id:          attempt.ID,
		ChallengeId: &attempt.ChallengeID,
		FeedbackId:  attempt.FeedbackID,
		Status:      domain.AttemptStatus(attempt.Status),
	}, nil
}

// ListAttemptsByChallenge retrieves all attempt entities from the repository
func (a *AttemptsRepository) ListAttemptsByUserChallenge(ctx context.Context, userId, challengeId int) ([]*domain.Attempt, error) {
	// Use the querier to get attempts from the database by challenge ID
	attempts, err := a.querier.ListAttemptsByUserChallenge(ctx, sqlc.ListAttemptsByUserChallengeParams{
		UserID: int64(userId),
		ID:     int64(challengeId),
	})
	if err != nil {
		return nil, err
	}

	// Map the retrieved attempts to the domain model
	var domainAttempts []*domain.Attempt
	for _, attempt := range attempts {
		domainAttempt := &domain.Attempt{
			Id:          attempt.ID,
			ChallengeId: &attempt.ChallengeID,
			FeedbackId:  attempt.FeedbackID,
			Status:      domain.AttemptStatus(attempt.Status),
		}
		domainAttempts = append(domainAttempts, domainAttempt)
	}

	return domainAttempts, nil
}
