package repository

import (
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Account interface {
	CreateAccountByUserId(userId int64) (int64, error)
	DeleteAccountByUserId(userId int64) error
	GetAccountByUserId(userId int64) (models.Account, error)
}

type Post interface {
	CreatePost(post models.Post) (int64, error)
}

type Repository struct {
	Account
	Post
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Account:NewAccountPostgres(db),
		Post:NewPostPostgres(db),
	}
}