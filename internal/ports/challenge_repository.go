package ports

import "github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"

// ChallengeRepository defines the interface for interacting with the DB
// to persist and retrieve Challenge entities
type ChallengeRepository interface {
	SaveChallenge(challenge *domain.Challenge) error
	GetChallengeByID(id string) (*domain.Challenge, error)
	UpdateChallenge(challenge *domain.Challenge) error
}
