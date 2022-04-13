package handler

import (
	"context"
	"fmt"
	"github.com/amrchnk/account_service/pkg/models"
	pb "github.com/amrchnk/account_service/proto"
	"time"
)

func (i *Implementation) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	images := make([]models.Image, 0, len(req.Post.Images))
	for i := range req.Post.Images {
		images = append(images, models.Image{
			Link: req.Post.Images[i].Link,
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

func (i *Implementation) DeletePostById (ctx context.Context,req *pb.DeletePostByIdRequest) (*pb.DeletePostByIdResponse, error) {
	err := i.Service.DeletePostById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePostByIdResponse{
		Message: fmt.Sprintf("Post with id=%d was deleted", req.Id),
	}, err
}
