package main

import (
	"net/http"

	"github.com/artyom-kalman/user-api-go/config"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config: %v", err)
		return
	}

	port, err := config.GetEnv("PORT")
	if err != nil {
		logger.Error("Failed to get PORT environment variable: %v", err)
		return
	}

	logger.Info("Starting server on port :%s", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Error("Failed to start server: %v", err)
		return
	}
}
