package models

import "time"

type Comment struct {
	Id        int64     `json:"id" db:"id"`
	Text      string    `json:"text" db:"text" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	PostId    int64     `json:"post_id" db:"post_id"`
	AccountId int64     `json:"-" db:"account_id"`
}