package models

import (
	"user_api/pkg/utils/errors"
)

type User struct {
	UserId     int    `json:"user_id"`
	Username   string `json:"username"`
	Balance    int    `json:"balance"`
	ReferrerId int    `json:"referrer_id"`
	CreatedAt  string `json:"created_at"`
}

type Users []User

type ReferrerRequest struct {
	ReferralID int `json:"referral_id"`
}

func (u *User) Validate() *errors.RestErr {
	if u.UserId <= 0 {
		return errors.NewBadRequestError("invalid user ID")
	}
	return nil
}
