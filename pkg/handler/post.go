package handler

import (
	"context"
	"github.com/amrchnk/account_service/pkg/models"
	pb "github.com/amrchnk/account_service/proto"
	"time"
)

func (i *Implementation) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	images := make([]models.Image, 0, len(req.Post.Images))
	for i := range req.Post.Images {
		images = append(images, models.Image{
			Link:   req.Post.Images[i].Link,
		})
	}

	post := models.Post{
		Title:       req.Post.Title,
		Description: req.Post.Description,
		CreatedAt:   time.Now(),
		Images:      images,
		AccountId:   req.Post.AccountId,
	}

	postId, err := i.Service.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{
		Id: postId,
	}, err
}
