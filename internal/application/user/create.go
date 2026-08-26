package user

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// CreateUserUseCase handles the creation of a new user in the application
type CreateUserUseCase struct {
	Logger         app.Logger
	UserRepository ports.UserRepository
}

// NewCreateUserUseCase creates a new instance of CreateUserUseCase
func NewCreateUserUseCase(logger app.Logger, userRepository ports.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		Logger:         logger,
		UserRepository: userRepository,
	}
}

// Execute creates a new user with the given userName and email, and saves it to the repository
func (c *CreateUserUseCase) Execute(ctx context.Context, userName, email string) (*domain.User, error) {
	// Create a new user entity
	user := &domain.User{
		UserName: userName,
		Email:    email,
	}

	// Save the user to the repository
	err := c.UserRepository.SaveUser(ctx, user)
	if err != nil {
		c.Logger.Error("failed to save user", "error", err)

		return nil, err
	}

	return user, nil
}
