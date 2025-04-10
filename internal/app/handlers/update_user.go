package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/artyom-kalman/user-api-go/internal/app/users"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func (h *UserHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request to update user")

	var newUserData users.User
	if err := json.NewDecoder(r.Body).Decode(&newUserData); err != nil {
		logger.Error("error decoding user data: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	logger.Debug("Received user data: %+v", newUserData)

	if err := h.repo.Update(&newUserData, r.Context()); err != nil {
		logger.Error("error updating user data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully updated user data")
}
