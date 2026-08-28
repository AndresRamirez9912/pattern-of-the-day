package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndresRamirez9912/pattern-of-the-day/internal/application/user"
)

// UserHandler only holds the use cases it actually needs
type UserHandler struct {
	createUser *user.CreateUserUseCase
	getUser    *user.GetUserUseCase
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(createUser *user.CreateUserUseCase, getUser *user.GetUserUseCase) *UserHandler {
	return &UserHandler{createUser: createUser, getUser: getUser}
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

// Get handles the request to retrieve a user by its ID.
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request URL.
	idStr := r.PathValue("user_id")
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert the user ID from string to integer.
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	// Execute the get user use case.
	found, err := h.getUser.Execute(r.Context(), userId)
	if err != nil {
		h.getUser.Logger.Error("error executing get user use case", "error", err.Error())

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved user details.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateUserResponse{
		Id:       found.Id,
		Username: found.UserName,
		Email:    found.Email,
	})
}
