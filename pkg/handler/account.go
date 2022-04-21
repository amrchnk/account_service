package handler

import (
	"context"
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
	log.Print("[INFO] User account was deleted")
	return &pb.DeleteAccountByUserIdResponse{Message: "Account was delete successful"}, nil
}

func (i *Implementation) GetAccountByUserId(ctx context.Context, req *pb.GetAccountByUserIdRequest) (*pb.GetAccountByUserIdResponse, error) {
	account, err := i.Service.GetAccountByUserId(req.UserId)
	if err != nil {
		return nil, err
	}

	respAccount := pb.Account{
		Id:        account.Id,
		UserId:    account.UserId,
		CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return &pb.GetAccountByUserIdResponse{
		Account: &respAccount,
	}, nil
}
