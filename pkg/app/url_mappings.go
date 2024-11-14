package app

import (
	"user_api/pkg/controllers"
	"user_api/pkg/middlewares"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.Use(middlewares.AuthMiddleware)

	router.GET("/users/:id/status", controllers.GetUserStatus)
	router.GET("/users/leaderboard", controllers.GetLeaderboard)
	router.POST("/users/:id/task/complete", controllers.CompleteTask)
	router.POST("/users/:id/referrer", controllers.AddReferrer)
}
