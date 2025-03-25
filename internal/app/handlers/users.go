package handlers

import (
	"net/http"

	"github.com/artyom-kalman/user-api-go/internal/app/repository"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: userRepo,
	}
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetUser(w, r)
	case http.MethodPost:
		h.handleNewUser(w, r)
	case http.MethodPut:
		h.handleUpdateUser(w, r)
	case http.MethodDelete:
		h.handleDeleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
