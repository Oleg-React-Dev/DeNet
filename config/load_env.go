package config

import (
	"user_api/pkg/logger"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		logger.Error("con not load environment variables:", err)
		return err
	}
	return nil
}
