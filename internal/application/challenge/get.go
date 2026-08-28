package challenge

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// GetChallengeUseCase retrieves a single challenge, with its clues, by ID.
type GetChallengeUseCase struct {
	Logger              ports.Logger
	ChallengeRepository ports.ChallengeRepository
}

// NewGetChallengeUseCase creates a new instance of GetChallengeUseCase
func NewGetChallengeUseCase(logger ports.Logger, challengeRepository ports.ChallengeRepository) *GetChallengeUseCase {
	return &GetChallengeUseCase{
		Logger:              logger,
		ChallengeRepository: challengeRepository,
	}
}

// Execute retrieves a challenge by its ID from the ChallengeRepository
func (g *GetChallengeUseCase) Execute(ctx context.Context, id int64) (*domain.Challenge, error) {
	challenge, err := g.ChallengeRepository.GetChallengeByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		g.Logger.Error("challenge not found", "challenge_id", id)
		return nil, domain.NewError("challenge not found", domain.ErrCodeNotFound)
	}
	if err != nil {
		g.Logger.Error("error fetching challenge", "challenge_id", id, "error", err.Error())
		return nil, err
	}

	return challenge, nil
}
