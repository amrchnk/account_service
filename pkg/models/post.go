package models

import "time"

type Post struct {
	Id          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Images      []Image   `json:"images"`
	Categories  []int64   `json:"categories"`
	AccountId   int64     `json:"account_id" db:"account_id" binding:"required"`
}

type Image struct {
	Id     int64  `json:"id" db:"id"`
	Link   string `json:"link" db:"link" binding:"required"`
	PostId int64  `json:"post_id" db:"post_id" binding:"required"`
}

type GetAllUsersPosts struct {
	Id          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Images      []string  `json:"images" db:"image"`
	Categories  []string   `json:"categories,omitempty" db:"category"`
	UserId     int64     `json:"user_id" db:"user_id" binding:"required"`
}
