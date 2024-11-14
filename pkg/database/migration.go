package database

import (
	"user_api/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	driver, err := postgres.WithInstance(Db, &postgres.Config{})
	if err != nil {
		logger.Error("could not create migration driver: %v", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		logger.Error("migration failed: %v", err)
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Error("could not apply migrations: %v", err)
		return err
	}

	logger.Info("Migrations applied successfully.")
	return nil
}
