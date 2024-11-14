package app

import (
	"os"
	"user_api/config"
	"user_api/pkg/database"
	"user_api/pkg/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	if err := config.LoadEnv(); err != nil {
		logger.Error("error occurred while loading env file:", err)
		panic(err)
	}

	if err := database.InitDB(); err != nil {
		logger.Error("error occurred while initializing database:", err)
		panic(err)
	}

	if err := database.RunMigrations(); err != nil {
		logger.Error("Migration failed: %v", err)
		panic(err)
	}

	mapUrls()

	logger.Info("about to start the application...")

	if err := router.Run(os.Getenv("PORT")); err != nil {
		logger.Error("error occurred while running http server:", err)
		panic(err)
	}

}
