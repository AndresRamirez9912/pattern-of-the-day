package rest

import (
	"net/http"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/adapters/inbound/rest/handlers"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/app"
)

// NewRouter registers every route it exposes
func NewRouter(services *app.Services) *http.ServeMux {
	// Create a new HTTP request router
	mux := http.NewServeMux()

	// Create the handlers
	userHandler := handlers.NewUserHandler(services.User.CreateUser, services.User.GetUser)
	challengeHandler := handlers.NewChallengeHandler(services.Challenge.CreateChallenge, services.Challenge.GetChallengeDetails)
	clueHandler := handlers.NewClueHandler(services.Clue.CreateClue, services.Challenge.GetChallenge)
	feedbackHandler := handlers.NewFeedbackHandler(services.Feedback.CreateFeedback, services.Attempts.GetAttempt, services.Challenge.GetChallenge)

	// Define the routes for the user handler
	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("GET /users/{user_id}", userHandler.Get)

	// Define the routes for the challenge handler
	mux.HandleFunc("GET /challenges/{challenge_id}", challengeHandler.Get)
	mux.HandleFunc("POST /challenges", challengeHandler.Create)

	// Define the routes for the clue handler
	mux.HandleFunc("POST /clues", clueHandler.Create)

	// Define the routes for the feedback handler
	mux.HandleFunc("POST /feedback", feedbackHandler.Create)

	// Define the route for the health check
	mux.HandleFunc("GET /health", HealthCheck)

	return mux
}
