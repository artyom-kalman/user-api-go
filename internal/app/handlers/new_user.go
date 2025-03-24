package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/artyom-kalman/user-api-go/internal/app"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received newUser request")

	var user app.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("error reading request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	logger.Debug("Received user: %+v", user)

	if user.Email == "" || user.Password == "" {
		logger.Error("error creating user: email or password is empty")
		http.Error(w, "email or password is empty", http.StatusBadRequest)
		return
	}

	if err := user.Save(r.Context()); err != nil {
		logger.Error("error saving user: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Successfully created user")
}
