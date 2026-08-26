package persistence

import (
	"context"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/persistence/sqlc"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Type assertion UserRepository must implement the UserRepository port
var _ ports.UserRepository = &UserRepository{}

type UserRepository struct {
	querier *sqlc.Queries
}

// NewUserRepository creates a new instance of UserRepository with the provided querier
func NewUserRepository(querier *sqlc.Queries) *UserRepository {
	return &UserRepository{
		querier: querier,
	}
}

// SaveUser saves a user entity to the repository
func (u *UserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	// Use the querier to create a new user in the database
	_, err := u.querier.CreateUser(ctx, sqlc.CreateUserParams{
		Username: user.UserName,
		Email:    user.Email,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID retrieves a user entity from the repository by its ID
func (u *UserRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	// Use the querier to get a user from the database by its ID
	user, err := u.querier.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Map the retrieved user to the domain model
	return &domain.User{
		Id:       user.ID,
		UserName: user.Username,
		Email:    user.Email,
	}, nil

}

// ListUserChallenges retrieves all challenges associated with a user
func (u *UserRepository) ListUserChallenges(ctx context.Context, userId int64) ([]*domain.Challenge, error) {
	challenges, err := u.querier.ListUserChallenges(ctx, userId)
	if err != nil {
		return nil, err
	}

	var result []*domain.Challenge
	for _, challenge := range challenges {
		// List clues associated to the challenge
		clues, err := u.querier.ListCluesByChallengeId(ctx, challenge.ID)
		if err != nil {
			return nil, err
		}

		var domainClues []domain.Clue
		for _, clue := range clues {
			domainClues = append(domainClues, domain.Clue{
				Id:            clue.ID,
				Description:   clue.Description,
				SequenceOrder: int(clue.SequenceOrder),
			})
		}

		// Append the challenge along with its associated clues to the result list
		result = append(result, &domain.Challenge{
			Id:          challenge.ID,
			Name:        challenge.Name,
			Description: challenge.Description,
			Dificulty:   domain.Difficulty(challenge.Difficulty),
			Type:        domain.ChallengeType(challenge.Type),
			Pattern:     domain.Pattern(challenge.TargetPattern),
			UserId:      challenge.UserID,
			Clues:       domainClues,
		})
	}

	return result, nil
}
