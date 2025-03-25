package main

import (
	"github.com/artyom-kalman/user-api-go/internal/app"
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

func main() {
	logger.Info("Starting app...")

	err := app.Start()
	if err != nil {
		logger.Error("error starting app: %v", err)
		return
	}
}
