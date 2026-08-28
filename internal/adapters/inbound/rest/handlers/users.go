package handlers

import (
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

}
