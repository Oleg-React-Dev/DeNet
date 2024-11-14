package models

import "user_api/pkg/utils/errors"

type TaskRequest struct {
	TaskId int `json:"task_id"`
}

func (t *TaskRequest) Validate() *errors.RestErr {
	if t.TaskId <= 0 {
		return errors.NewBadRequestError("invalid task ID")
	}
	return nil
}
