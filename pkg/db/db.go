package db

import (
	"database/sql"

	"github.com/artyom-kalman/user-api-go/config"
)

type Database struct {
	conn *sql.DB
}

var openedConn *Database

func GetDatabase() *Database {
	if openedConn == nil {
		databaseConfig, _ := config.GetDBConfig()
		Connect(databaseConfig)
	}

	return openedConn
}
