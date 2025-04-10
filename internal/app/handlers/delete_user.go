package handlers

import (
	"net/http"

	"github.com/artyom-kalman/user-api-go/internal/app/repository"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func (h *UserHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request to delete user")

	userId := r.URL.Query().Get("id")
	if userId == "" {
		logger.Error("User ID is missing")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	logger.Info("Deleting user with ID = %s", userId)

	user, err := h.repo.GetUserById(userId, r.Context())
	if err != nil {
		if err == repository.ErrUserNotFound {
			logger.Info("User with id = %s", userId)
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			logger.Error("error getting user with id = %s: %v", userId, err)
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		}
		return
	}

	err = h.repo.Delete(user, r.Context())
	if err != nil {
		logger.Error("error deleting user with id = %s: %v", userId, err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully deleted user")
}
