package handler

import (
	"context"
	"fmt"
	"github.com/amrchnk/account_service/pkg/models"
	pb "github.com/amrchnk/account_service/proto"
	"log"
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
		UpdatedAt:   time.Now(),
		Images:      images,
		Categories:  req.Post.Categories,
		AccountId:   req.Post.AccountId,
	}

	postId, err := i.Service.CreatePost(post)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	log.Printf("[INFO] Post with id = %d was created", postId)
	return &pb.CreatePostResponse{
		Id: postId,
	}, err
}

func (i *Implementation) GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error) {
	resp, err := i.Service.GetPostById(req.Id)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	return &pb.GetPostByIdResponse{
		Post: &pb.PostV2{
			Id:          resp.Id,
			Title:       resp.Title,
			Description: resp.Description,
			CreatedAt:   resp.CreatedAt.Format("2006-01-02 15:04:05"),
			Images:      resp.Images,
			Categories:  resp.Categories,
			UpdatedAt:   resp.UpdatedAt.Format("2006-01-02 15:04:05"),
			UserId:      resp.UserId,
		},
	}, err
}

func (i *Implementation) UpdatePostById(ctx context.Context, req *pb.UpdatePostByIdRequest) (*pb.UpdatePostByIdResponse, error) {

	request := models.UpdatePost{
		Id:          req.PostId,
		Title:       req.Title,
		Description: req.Description,
		Images:      req.Images,
		Categories:  req.Categories,
	}

	resp, err := i.Service.UpdatePostByd(request)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	return &pb.UpdatePostByIdResponse{
		Message: resp,
	}, err
}

func (i *Implementation) DeletePostById(ctx context.Context, req *pb.DeletePostByIdRequest) (*pb.DeletePostByIdResponse, error) {
	err := i.Service.DeletePostById(req.Id)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}
	return &pb.DeletePostByIdResponse{
		Message: fmt.Sprintf("Post with id=%d was deleted", req.Id),
	}, err
}

func (i *Implementation) GetPostsByUserId(ctx context.Context, req *pb.GetUserPostsRequest) (*pb.GetUserPostsResponse, error) {
	posts, err := i.Service.GetPostsByUserId(req.UserId)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	postsResp := make([]*pb.Post, 0, len(posts))

	for i := range posts {
		images := make([]*pb.Image, 0, len(posts[i].Images))
		if len(posts[i].Images) != 0 {
			for _, image := range posts[i].Images {
				images = append(images, &pb.Image{
					Id:     image.Id,
					Link:   image.Link,
					PostId: image.PostId,
				})
			}
		}

		postsResp = append(postsResp, &pb.Post{
			Id:          posts[i].Id,
			Title:       posts[i].Title,
			Description: posts[i].Description,
			CreatedAt:   posts[i].CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   posts[i].UpdatedAt.Format("2006-01-02 15:04:05"),
			Categories:  posts[i].Categories,
			Images:      images,
			AccountId:   posts[i].AccountId,
		})
	}
	return &pb.GetUserPostsResponse{
		UserId: req.UserId,
		Posts:  postsResp,
	}, err
}

func (i *Implementation) GetAllUsersPosts(ctx context.Context, req *pb.GetAllUsersPostsRequest) (*pb.GetAllUsersPostsResponse, error) {
	posts, err := i.Service.GetAllUsersPosts(req.Offset, req.Limit, req.Sorting)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	postsResp := make([]*pb.PostV2, 0, len(posts))
	if len(posts) > 0 {
		for _, post := range posts {
			postsResp = append(postsResp, &pb.PostV2{
				Id:          post.Id,
				Title:       post.Title,
				Description: post.Description,
				CreatedAt:   post.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   post.UpdatedAt.Format("2006-01-02 15:04:05"),
				Images:      post.Images,
				Categories:  post.Categories,
				UserId:      post.UserId,
			})
		}
	}

	return &pb.GetAllUsersPostsResponse{
		Posts: postsResp,
	}, nil
}
