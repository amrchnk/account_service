package models

import "time"

type Account struct {
	Id        int64     `json:"id" db:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created-at" db:"created_at"`
}

type Comment struct {
	Id        int64     `json:"id" db:"id"`
	Text      string    `json:"text" db:"text" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	PostId    int64     `json:"post_id" db:"post_id"`
	AccountId int64     `json:"account_id" db:"account_id"`
}

type Category struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"category" db:"category"`
}

type Tag struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"tag" db:"tag"`
}
