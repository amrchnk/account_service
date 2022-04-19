package service

import (
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/amrchnk/account_service/pkg/repository"
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
}

type Service struct {
	Account
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
		Post:    NewPostService(repos.Post),
	}
}
