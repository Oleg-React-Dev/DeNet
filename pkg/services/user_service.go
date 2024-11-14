package services

import (
	"user_api/pkg/models"
	"user_api/pkg/utils/errors"
)

func GetUserStatus(id int) (*models.User, *errors.RestErr) {
	dao := &models.User{
		UserId: id,
	}

	if err := dao.Validate(); err != nil {
		return nil, err
	}

	if err := dao.FinedById(); err != nil {
		return nil, err
	}

	return dao, nil
}

func GetLeaderboard() (*models.Users, *errors.RestErr) {
	var dao models.User
	return dao.GetLeaderboard()
}

func CompleteTask(userId int, task models.TaskRequest) *errors.RestErr {
	userDao := &models.User{
		UserId: userId,
	}

	if err := userDao.Validate(); err != nil {
		return err
	}

	if err := task.Validate(); err != nil {
		return err
	}

	if err := userDao.CompleteTask(task.TaskId); err != nil {
		return err
	}

	return nil
}

func AddReferrer(userId, referralId int) *errors.RestErr {
	userDao := &models.User{
		UserId:     userId,
		ReferrerId: referralId,
	}

	if err := userDao.Validate(); err != nil {
		return err
	}

	if err := userDao.AddReferrer(); err != nil {
		return err
	}

	return nil
}
