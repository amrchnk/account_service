package models

import "time"

type Account struct {
	Id           int64     `json:"id" db:"id"`
	UserId       int64     `json:"user_id" db:"user_id"`
	ProfileImage string    `json:"profile_image" db:"profile_image"`
	CreatedAt    time.Time `json:"created-at" db:"created_at"`
}

type UpdateAccountInfo struct {
	UserId       int64  `json:"user_id" db:"user_id"`
	ProfileImage string `json:"profile_image" db:"profile_image"`
}
