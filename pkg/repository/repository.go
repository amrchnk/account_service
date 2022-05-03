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
	DeletePostById(postId int64) error
	GetPostById(postId int64) (models.Post, error)
	GetPostsByUserId(userId int64) ([]models.Post, error)
	UpdatePostByd(post models.Post) (string, error)
	GetAllUsersPosts(offset, limit int64, sorting string) ([]models.GetAllUsersPosts, error)
}

type Images interface {
	GetImagesFromPost(postId int64) ([]models.Image, error)
}

type Repository struct {
	Account
	Post
	Images
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Account: NewAccountPostgres(db),
		Post:    NewPostPostgres(db),
		Images:  NewImagesPostgres(db),
	}
}
