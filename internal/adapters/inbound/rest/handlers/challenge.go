package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/challenge"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/domain"
	"github.com/AndresRamirez9912/pattern-of-the-day/internal/ports"
)

// ChallengeHandler only holds the use cases it actually needs
type ChallengeHandler struct {
	createChallenge *challenge.CreateChallengeUseCase
	getChallenge    *challenge.GetChallengeUseCase
}

// NewChallengeHandler creates a new ChallengeHandler.
func NewChallengeHandler(
	createChallenge *challenge.CreateChallengeUseCase,
	getChallenge *challenge.GetChallengeUseCase,
) *ChallengeHandler {
	return &ChallengeHandler{
		createChallenge: createChallenge,
		getChallenge:    getChallenge,
	}
}

// Create handles the request to create a new challenge.
func (h *ChallengeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateChallengeRequest

	// Decode the JSON request body.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.createChallenge.Logger.Error("error decoding request body", "error", err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the create challenge use case.
	genReqPayload := ports.ChallengeGenerationRequest{
		Topic:      req.Topic,
		Difficulty: domain.Difficulty(req.Difficulty),
		Target:     req.Target,
		Type:       domain.ChallengeType(req.Type),
	}

	challenge, attempt, err := h.createChallenge.Execute(r.Context(), req.UserName, genReqPayload, ".")
	if err != nil {
		h.createChallenge.Logger.Error("error executing create challenge use case", "error", err.Error())

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.createChallenge.Logger.Info("successfully created challenge", "challenge_id", challenge.Id)

	// Create response
	resp := CreateChallengeResponse{
		Challenge: Challenge{
			ChallengeId: challenge.Id,
			Name:        challenge.Name,
			Description: challenge.Description,
			Difficulty:  string(challenge.Difficulty),
			Type:        string(challenge.Type),
			Target:      challenge.Target,
		},
		Attempt: Attempt{
			AttemptId: attempt.Id,
			Status:    string(attempt.Status),
			Sequence:  attempt.SequenceOrder,
		},
	}

	// Respond with the created challenge details.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Get handles the request to retrieve a challenge by its ID.
func (h *ChallengeHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Extract the challenge ID from the request URL.
	idStr := r.PathValue("challenge_id")
	if idStr == "" {
		h.getChallenge.Logger.Error("challenge ID is required")

		http.Error(w, "Challenge ID is required", http.StatusBadRequest)
		return
	}

	// Convert the challenge ID from string to integer.
	challengeID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.getChallenge.Logger.Error("invalid challenge ID format", "error", err.Error())

		http.Error(w, "Invalid Challenge ID format", http.StatusBadRequest)
		return
	}

	// Execute the get challenge use case.
	challenge, err := h.getChallenge.Execute(r.Context(), challengeID)
	if err != nil {
		h.getChallenge.Logger.Error("error executing get challenge use case", "error", err.Error())

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create response
	resp := CreateChallengeResponse{
		Challenge: Challenge{
			ChallengeId: challenge.Id,
			Name:        challenge.Name,
			Description: challenge.Description,
			Difficulty:  string(challenge.Difficulty),
			Type:        string(challenge.Type),
			Target:      challenge.Target,
		},
	}

	// Respond with the retrieved challenge details.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
