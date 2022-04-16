package service

import (
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/amrchnk/account_service/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post models.Post) (int64, error) {
	return s.repo.CreatePost(post)
}

func (s *PostService) DeletePostById(postId int64) error{
	return s.repo.DeletePostById(postId)
}

func (s *PostService) GetPostById(postId int64) (models.Post, error){
	return s.repo.GetPostById(postId)
}

func (s *PostService) GetPostsByUserId(userId int64) ([]models.Post, error){
	return s.repo.GetPostsByUserId(userId)
}