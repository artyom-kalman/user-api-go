package db

import (
	"fmt"

	"github.com/artyom-kalman/user-api-go/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration() error {
	logger.Info("Running migration...")

	if openedConn == nil {
		return fmt.Errorf("error running migration: database connection is closed")
	}

	driver, err := postgres.WithInstance(openedConn.conn, &postgres.Config{})
	if err != nil {
		return err
	}
	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations/",
		"user_api",
		driver,
	)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err == migrate.ErrNoChange {
		logger.Info("Database schema is up to date")
	} else if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("Successfully migrated database schema")
	}

	return nil
}
