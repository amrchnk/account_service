package models

import "time"

type Account struct {
	Id        int64     `json:"id" db:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Category struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"category" db:"category"`
}

type Tag struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"tag" db:"tag"`
}
