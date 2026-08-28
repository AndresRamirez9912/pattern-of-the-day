package handlers

import (
	"net/http"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/challenge"
)

// ChallengeHandler only holds the use cases it actually needs
type ChallengeHandler struct {
	createChallenge *challenge.CreateChallengeUseCase
}

// NewChallengeHandler creates a new ChallengeHandler.
func NewChallengeHandler(createChallenge *challenge.CreateChallengeUseCase) *ChallengeHandler {
	return &ChallengeHandler{createChallenge: createChallenge}
}

// Create handles the request to create a new challenge.
func (h *ChallengeHandler) Create(w http.ResponseWriter, r *http.Request) {}
