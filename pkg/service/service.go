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
	GetPostById(postId int64) (models.PostV2, error)
	UpdatePostByd(post models.UpdatePost) (string, error)
	GetPostsByUserId(userId int64) ([]models.Post, error)
	GetAllUsersPosts(offset, limit int64, sorting string) ([]models.PostV2, error)
}

/*type Comments interface{
	CreateComment(post models.Comment) (int64, error)
	DeleteCommentById(CommentId int64) error
	GetPostComments(postId int64) ([]models.Comment, error)
}
*/
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
