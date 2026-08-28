package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// GetUserUseCase retrieves a user by ID.
type GetUserUseCase struct {
	Logger         ports.Logger
	UserRepository ports.UserRepository
}

// NewGetUserUseCase creates a new GetUserUseCase.
func NewGetUserUseCase(logger ports.Logger, userRepository ports.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		Logger:         logger,
		UserRepository: userRepository,
	}
}

// Execute retrieves a user by its ID.
func (g *GetUserUseCase) Execute(ctx context.Context, id int64) (*domain.User, error) {
	found, err := g.UserRepository.GetUserByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		g.Logger.Error("user not found", "user_id", id)
		return nil, domain.NewError("user not found", domain.ErrCodeNotFound)
	}
	if err != nil {
		g.Logger.Error("error fetching user", "user_id", id, "error", err.Error())
		return nil, err
	}

	return found, nil
}
