package main

import (
	"net/http"

	"github.com/artyom-kalman/user-api-go/config"
	"github.com/artyom-kalman/user-api-go/pkg/db"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config: %v", err)
		return
	}

	dbConfig, err := config.GetDBConfig()
	if err != nil {
		logger.Error("failed to get db config: %v", err)
		return
	}

	err = db.Connect(dbConfig)
	if err != nil {
		logger.Error("failed to connect to database: %v", err)
		return
	}
	defer db.Close()

	err = db.RunMigration()
	if err != nil {
		logger.Error("failed to migrate database: %v", err)
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
