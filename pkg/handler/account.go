package handler

import (
	"context"
	"fmt"
	"github.com/amrchnk/account_service/pkg/models"
	pb "github.com/amrchnk/account_service/proto"
	"log"
)

func (i *Implementation) CreateAccountByUserId(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	accountId, err := i.Service.CreateAccountByUserId(req.UserId)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Account with id = %d was created", accountId)
	return &pb.CreateAccountResponse{
		AccountId: accountId,
	}, err
}

func (i *Implementation) DeleteAccountByUserId(ctx context.Context, req *pb.DeleteAccountByUserIdRequest) (*pb.DeleteAccountByUserIdResponse, error) {
	err := i.Service.DeleteAccountByUserId(req.UserId)
	if err != nil {
		return &pb.DeleteAccountByUserIdResponse{Message: string(err.Error())}, err
	}
	log.Println(fmt.Sprintf("[INFO] Account with userId = %d was delete successful",req.UserId))
	return &pb.DeleteAccountByUserIdResponse{Message: fmt.Sprintf("Account with userId = %d was delete successful",req.UserId)}, nil
}

func (i *Implementation) GetAccountByUserId(ctx context.Context, req *pb.GetAccountByUserIdRequest) (*pb.GetAccountByUserIdResponse, error) {
	account, err := i.Service.GetAccountByUserId(req.UserId)
	if err != nil {
		return nil, err
	}

	respAccount := pb.Account{
		Id:        account.Id,
		UserId:    account.UserId,
		ProfileImage: account.ProfileImage,
		CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return &pb.GetAccountByUserIdResponse{
		Account: &respAccount,
	}, nil
}

func (i *Implementation) UpdateAccountByUserId(ctx context.Context, req *pb.UpdateAccountByUserIdRequest) (*pb.UpdateAccountByUserIdResponse, error) {
	updateReq := models.UpdateAccountInfo{
		UserId:       req.NewInfo.UserId,
		ProfileImage: req.NewInfo.ProfileImage,
	}
	accountId, err := i.Service.UpdateAccountInfo(updateReq)
	if err != nil {
		return &pb.UpdateAccountByUserIdResponse{
			Message: err.Error(),
		}, err
	}

	return &pb.UpdateAccountByUserIdResponse{
		Message: fmt.Sprintf("Account with id = %d updated successfully", accountId),
	}, err
}
