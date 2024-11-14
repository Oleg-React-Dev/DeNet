package controllers

import (
	"net/http"
	"strconv"
	"user_api/pkg/models"
	"user_api/pkg/services"
	"user_api/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

func GetUserStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid user ID")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, statusErr := services.GetUserStatus(id)
	if err != nil {
		c.JSON(statusErr.Status, statusErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetLeaderboard(c *gin.Context) {
	leaderboard, err := services.GetLeaderboard()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, leaderboard)
}

func CompleteTask(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid user ID")
		c.JSON(restErr.Status, restErr)
		return
	}

	var task models.TaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		restErr := errors.NewBadRequestError("Invalid input")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := services.CompleteTask(userId, task); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Task completed"})
}

func AddReferrer(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid user ID")
		c.JSON(restErr.Status, restErr)
		return
	}

	var referrer models.ReferrerRequest
	if err := c.ShouldBindJSON(&referrer); err != nil {
		restErr := errors.NewBadRequestError("Invalid input")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := services.AddReferrer(user_id, referrer.ReferralID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Referrer added"})
}
