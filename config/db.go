package config

import (
	"github.com/artyom-kalman/user-api-go/pkg/logger"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDBConfig() (*DBConfig, error) {
	host, err := GetEnv("POSTGRES_HOST")
	if err != nil {
		return nil, err
	}

	port, err := GetEnv("POSTGRES_PORT")
	if err != nil {
		return nil, err
	}

	user, err := GetEnv("POSTGRES_USER")
	if err != nil {
		return nil, err
	}

	password, err := GetEnv("POSTGRES_PASSWORD")
	if err != nil {
		return nil, err
	}

	dbName, err := GetEnv("POSTGRES_DB")
	if err != nil {
		return nil, err
	}

	logger.Logger.Info("Loaded database configuration")

	return &DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}, nil
}
