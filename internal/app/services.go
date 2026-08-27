package app

import (
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/llm/ollama"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/persistence"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/persistence/sqlc"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app/config"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/attempt"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/challenge"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/clue"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/feedback"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/user"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// Services represents the services layer of the application
type Services struct {
	User      *UserUseCases
	Challenge *ChallengeUseCases
	Clue      *ClueUseCases
	Attempts  *AttemptUseCases
	Feedback  *FeedbackUseCases
}

// UserUseCases represents the use cases related to user management in the application.
type UserUseCases struct {
	CreateUser *user.CreateUserUseCase
}

// ChallengeUseCases represents the use cases related to challenge management in the application.
type ChallengeUseCases struct {
	CreateChallenge *challenge.CreateChallengeUseCase
	GetChallenge    *challenge.GetChallengeUseCase
}

// ClueUseCases represents the use cases related to clue management in the application.
type ClueUseCases struct {
	CreateClue *clue.CreateClueUseCase
}

// AttemptUseCases represents the use cases related to attempt management in the application.
type AttemptUseCases struct {
	CreateAttempt *attempt.CreateAttemptUseCase
	GetAttempt    *attempt.GetAttemptUseCase
}

// FeedbackUseCases represents the use cases related to feedback management in the application.
type FeedbackUseCases struct {
	CreateFeedback *feedback.CreateFeedbackUseCase
}

// newServices creates a new instance of Services with the required dependencies
func (a *App) newServices(cfg *config.AppConfig, logger *Logger) *Services {
	// Create queries
	userQuery := sqlc.New(a.db)

	// Initialize repositories
	userRepository := persistence.NewUserRepository(userQuery)
	challengeRepository := persistence.NewChallengeRepository(userQuery)
	clueRepository := persistence.NewClueRepository(userQuery)
	attemptRepository := persistence.NewAttemptsRepository(userQuery)
	feedbackRepository := persistence.NewFeedbackRepository(userQuery)

	// Initialize LLM Provider
	ollamaClient := ollama.NewClient(cfg.Llm.BaseUrl, &cfg.Llm.ApiKey)
	ollamaProvider := ollama.NewProvider(ollamaClient, cfg.Llm.Model, logger)

	// Initialize use cases
	userUseCase := a.createUser(logger, userRepository)
	challengeUseCase := a.createChallenge(logger, ollamaProvider, challengeRepository, attemptRepository, userRepository)
	clueUseCase := a.createClue(logger, ollamaProvider, clueRepository)
	attemptsUseCase := a.createAttempt(logger, attemptRepository)
	feedbackUseCase := a.createFeedback(logger, feedbackRepository, ollamaProvider, attemptRepository)

	return &Services{
		User:      userUseCase,
		Challenge: challengeUseCase,
		Clue:      clueUseCase,
		Attempts:  attemptsUseCase,
		Feedback:  feedbackUseCase,
	}
}

// createUser initializes the UserUseCases with the required dependencies.
func (a *App) createUser(logger ports.Logger, userRepository ports.UserRepository) *UserUseCases {
	createUser := user.NewCreateUserUseCase(logger, userRepository)

	return &UserUseCases{
		CreateUser: createUser,
	}
}

// createChallenge initializes the ChallengeUseCases with the required dependencies.
func (a *App) createChallenge(
	logger ports.Logger,
	ollamaProvider ports.LLMProvider,
	challengeRepository ports.ChallengeRepository,
	attemptsRepository ports.AttemptsRepository,
	userRepository ports.UserRepository,
) *ChallengeUseCases {
	createChallenge := challenge.NewCreateChallengeUseCase(logger, ollamaProvider, challengeRepository, attemptsRepository, userRepository)
	getChallenge := challenge.NewGetChallengeUseCase(logger, challengeRepository)

	return &ChallengeUseCases{
		CreateChallenge: createChallenge,
		GetChallenge:    getChallenge,
	}
}

// createClue initializes the ClueUseCases with the required dependencies.
func (a *App) createClue(logger ports.Logger, ollamaProvider ports.LLMProvider, clueRepository ports.ClueRepository) *ClueUseCases {
	createClue := clue.NewCreateClueUseCase(logger, ollamaProvider, clueRepository)

	return &ClueUseCases{
		CreateClue: createClue,
	}
}

// createAttempt initializes the AttemptUseCases with the required dependencies.
func (a *App) createAttempt(logger ports.Logger, attemptRepository ports.AttemptsRepository) *AttemptUseCases {
	createAttempt := attempt.NewCreateAttemptUseCase(logger, attemptRepository)
	getAttempt := attempt.NewGetAttemptUseCase(logger, attemptRepository)

	return &AttemptUseCases{
		CreateAttempt: createAttempt,
		GetAttempt:    getAttempt,
	}
}

// createFeedback initializes the FeedbackUseCases with the required dependencies.
func (a *App) createFeedback(
	logger ports.Logger,
	feedbackRepository ports.FeedbackRepository,
	ollamaProvider ports.LLMProvider,
	attemptRepository ports.AttemptsRepository,
) *FeedbackUseCases {
	createFeedback := feedback.NewCreateFeedbackUseCase(logger, feedbackRepository, ollamaProvider, attemptRepository)

	return &FeedbackUseCases{
		CreateFeedback: createFeedback,
	}
}
