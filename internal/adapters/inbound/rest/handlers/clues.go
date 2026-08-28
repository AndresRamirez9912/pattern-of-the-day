package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/challenge"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/clue"
)

// ClueHandler handles HTTP requests related to clues.
type ClueHandler struct {
	createClue   *clue.CreateClueUseCase
	getChallenge *challenge.GetChallengeUseCase
}

// NewClueHandler creates a new instance of ClueHandler.
func NewClueHandler(
	createClue *clue.CreateClueUseCase,
	getChallenge *challenge.GetChallengeUseCase,
) *ClueHandler {
	return &ClueHandler{
		createClue:   createClue,
		getChallenge: getChallenge,
	}
}

// Create handles the request to create a new clue.
func (h *ClueHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a CreateClueRequest struct
	var req CreateClueRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.createClue.Logger.Error("failed to decode create clue request", "error", err)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate that the challenge exists before creating a clue
	challenge, err := h.getChallenge.Execute(r.Context(), req.ChallengeId)
	if err != nil {
		h.getChallenge.Logger.Error("failed to get challenge", "error", err)

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Call the use case to create a new clue
	clue, err := h.createClue.Execute(r.Context(), challenge, ".")
	if err != nil {
		h.createClue.Logger.Error("failed to create clue", "error", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created clue as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		CreateClueResponse{
			Clue: Clue{
				ClueId:      clue.Id,
				ChallengeId: challenge.Id,
				Description: clue.Description,
			},
		},
	)
}
