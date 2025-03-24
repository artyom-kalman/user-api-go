package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/artyom-kalman/user-api-go/internal/app"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Request getUser request")

	var userID = r.URL.Query().Get("id")
	if userID == "" {
		logger.Error("User ID is empty")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	logger.Debug("Received user ID: %s", userID)

	user, err := app.GetUserById(r.Context(), userID)
	if err != nil {
		if err == app.ErrUserNotFound {
			logger.Error("User not found")
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			logger.Error("error finding user with id = %s: %v", userID, err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		logger.Error("error encoding user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully retrieved user")
}
