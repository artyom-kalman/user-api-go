package db

import (
	"database/sql"

	"github.com/artyom-kalman/user-api-go/config"
)

type Database struct {
	conn *sql.DB
}

var databaseConnection *Database

func GetDatabase() *Database {
	if databaseConnection == nil {
		databaseConfig, _ := config.GetDBConfig()
		ConnectToDatabase(databaseConfig)
	}

	return databaseConnection
}
