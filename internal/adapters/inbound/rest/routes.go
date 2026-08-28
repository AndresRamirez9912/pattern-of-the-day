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
	userHandler := handlers.NewUserHandler(services.User.CreateUser)
	challengeHandler := handlers.NewChallengeHandler(services.Challenge.CreateChallenge, services.Challenge.GetChallenge)
	clueHandler := handlers.NewClueHandler(services.Clue.CreateClue, services.Challenge.GetChallenge)

	// Define the routes for the user handler
	mux.HandleFunc("POST /users", userHandler.Create)

	// Define the routes for the challenge handler
	mux.HandleFunc("GET /challenges/{challenge_id}", challengeHandler.Get)
	mux.HandleFunc("POST /challenges", challengeHandler.Create)

	// Define the routes for the clue handler
	mux.HandleFunc("POST /clues", clueHandler.Create)

	// Define the route for the health check
	mux.HandleFunc("GET /health", HealthCheck)

	return mux
}
