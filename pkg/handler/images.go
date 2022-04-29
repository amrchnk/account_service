package handler

import (
	"context"
	pb "github.com/amrchnk/account_service/proto"
	"log"
)

func (i *Implementation) GetImagesFromPost(ctx context.Context, req *pb.GetImagesFromPostRequest) (*pb.GetImagesFromPostResponse, error) {
	images, err := i.Service.GetImagesFromPost(req.PostId)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return &pb.GetImagesFromPostResponse{}, err
	}
	imagesResp := make([]*pb.Image, 0, len(images))
	for index := range images {
		imagesResp = append(imagesResp, &pb.Image{
			Id:     images[index].Id,
			Link:   images[index].Link,
			PostId: images[index].PostId,
		})
	}

	return &pb.GetImagesFromPostResponse{
		Images: imagesResp,
	}, nil
}
