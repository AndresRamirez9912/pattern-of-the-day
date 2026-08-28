package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/user"
)

// UserHandler only holds the use cases it actually needs
type UserHandler struct {
	createUser *user.CreateUserUseCase
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(createUser *user.CreateUserUseCase) *UserHandler {
	return &UserHandler{createUser: createUser}
}

// Create handles the HTTP request for creating a new user.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	// Decode the JSON request body.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.createUser.Logger.Error("error decoding request body", "error", err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the create user use case.
	resp, err := h.createUser.Execute(r.Context(), req.Username, req.Email)
	if err != nil {
		h.createUser.Logger.Error("error executing create user use case", "error", err.Error())

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created user details.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateUserResponse{
		Id:       resp.Id,
		Username: resp.UserName,
		Email:    resp.Email,
	})
}
