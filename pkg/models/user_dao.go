package models

import (
	"database/sql"
	"strings"
	"user_api/pkg/database"
	"user_api/pkg/logger"
	"user_api/pkg/utils/errors"
)

const (
	queryGetUserById    = "SELECT username, balance, referrer_id, created_at FROM users WHERE user_id=$1;"
	queryGetLeaderboard = "SELECT user_id, username, balance, referrer_id, created_at FROM users ORDER BY balance DESC LIMIT 10;"
	queryCompleteTask   = "INSERT INTO user_tasks (user_id, task_id) VALUES ($1, $2) ON CONFLICT (user_id, task_id) DO NOTHING RETURNING user_id;"
	queryAddReferrer    = "UPDATE users SET referrer_id = $1 WHERE user_id = $2 RETURNING user_id;"

	duplicateErr = "duplicate key value violates unique constraint"
)

func (u *User) FinedById() *errors.RestErr {
	row := database.Db.QueryRow(queryGetUserById, u.UserId)
	scanErr := row.Scan(&u.Username, &u.Balance, &u.ReferrerId, &u.CreatedAt)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return errors.NewNotFoundError("invalid user credentials")
		}

		logger.Error("error when trying to get user by id", scanErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (u *User) GetLeaderboard() (*Users, *errors.RestErr) {
	rows, err := database.Db.Query(queryGetLeaderboard)
	if err != nil {
		logger.Error("error when trying to get leaderboard", err)
		return nil, errors.NewInternalServerError("database error")
	}

	results := make(Users, 0, 10)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserId, &u.Username, &u.Balance, &u.ReferrerId, &u.CreatedAt); err != nil {
			logger.Error("error when trying to scan row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("no users found")
	}
	return &results, nil
}

func (u *User) CompleteTask(taskId int) *errors.RestErr {
	saveErr := database.Db.QueryRow(queryCompleteTask, u.UserId, taskId).Scan(&u.UserId)

	if saveErr != nil {
		if strings.Contains(saveErr.Error(), duplicateErr) {
			return errors.NewBadRequestError("the task already completed for given user")
		}
		logger.Error("error when trying to save completed task", saveErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (u *User) AddReferrer() *errors.RestErr {
	saveErr := database.Db.QueryRow(queryAddReferrer, u.UserId, u.ReferrerId).Scan(&u.UserId)

	if saveErr != nil {
		if strings.Contains(saveErr.Error(), duplicateErr) {
			return errors.NewBadRequestError("the referrer already exists for given user")
		}
		logger.Error("error when trying to add referrer", saveErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
