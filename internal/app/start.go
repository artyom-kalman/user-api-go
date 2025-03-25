package app

import (
	"net/http"

	"github.com/artyom-kalman/user-api-go/config"
	"github.com/artyom-kalman/user-api-go/internal/app/handlers"
	"github.com/artyom-kalman/user-api-go/internal/app/repository"
	"github.com/artyom-kalman/user-api-go/pkg/db"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func Start() error {
	err := config.LoadConfig()
	if err != nil {
		logger.Error("error loading config: %v", err)
		return err
	}

	dbConfig, err := config.GetDBConfig()
	if err != nil {
		logger.Error("error getting db config: %v", err)
		return err
	}

	err = db.Connect(dbConfig)
	if err != nil {
		logger.Error("error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	err = db.RunMigration()
	if err != nil {
		logger.Error("error migrating database: %v", err)
		return err
	}

	userRepo := repository.NewUserRepository(db.GetDatabase())
	handler := handlers.NewUserHandler(userRepo)

	http.HandleFunc("/users", handler.HandleUsers)

	port, err := config.GetEnv("PORT")
	if err != nil {
		logger.Error("error getting PORT environment variable: %v", err)
		return err
	}

	logger.Info("Starting server on port :%s", port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Error("error starting server: %v", err)
		return err
	}

	return nil
}
