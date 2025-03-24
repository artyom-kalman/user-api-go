package config

import (
	"fmt"
	"os"

	"github.com/artyom-kalman/user-api-go/pkg/logger"
	"github.com/joho/godotenv"
)

func LoadConfig() error {
	if _, err := os.Stat(".env"); err != nil {
		return fmt.Errorf(".env file not found")
	}

	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	logger.Info("Successfully loaded config")

	return nil
}

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return value, nil
}
