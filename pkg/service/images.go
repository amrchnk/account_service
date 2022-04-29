package service

import (
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/amrchnk/account_service/pkg/repository"
)

type ImagesService struct {
	repo repository.Images
}

func NewImagesService(repo repository.Images) *ImagesService {
	return &ImagesService{repo: repo}
}

func (s *ImagesService) GetImagesFromPost(postId int64) ([]models.Image, error) {
	return s.repo.GetImagesFromPost(postId)
}
