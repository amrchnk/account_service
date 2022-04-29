package service

import (
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/amrchnk/account_service/pkg/repository"
)

type Account interface {
	CreateAccountByUserId(userId int64) (int64, error)
	DeleteAccountByUserId(userId int64) error
	GetAccountByUserId(userId int64) (models.Account, error)
	UpdateAccountInfo(info models.UpdateAccountInfo) (int64, error)
}

type Post interface {
	CreatePost(post models.Post) (int64, error)
	DeletePostById(postId int64) error
	GetPostById(postId int64) (models.Post, error)
	UpdatePostByd(post models.Post) (string, error)
	GetPostsByUserId(userId int64) ([]models.Post, error)
}

type Images interface {
	GetImagesFromPost(postId int64) ([]models.Image, error)
}

type Service struct {
	Account
	Post
	Images
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repos.Account),
		Post:    NewPostService(repos.Post),
		Images:  NewImagesService(repos.Images),
	}
}
