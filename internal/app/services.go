package app

import "github.com/AndresRamirez9912/pattern-of-the-day/internal/app/config"

// Services represents the services layer of the application
type Services struct {
	// Challenges *ChallengeService
	// Clues      *ClueService
	// Feedback   *FeedbackService
}

// NewServices creates a new instance of Services with the required dependencies
func NewServices(cfg *config.AppConfig, logger *Logger) *Services {
	// Initialize each service
	return &Services{
		// Challenges: NewChallengeService(),
		// Clues:      NewClueService(),
		// Feedback:   NewFeedbackService(),
	}
}
