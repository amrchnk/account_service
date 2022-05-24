package service

import (
	"database/sql"
	"errors"
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/amrchnk/account_service/pkg/repository"
	"log"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccountByUserId(userId int64) (int64, error) {
	return s.repo.CreateAccountByUserId(userId)
}

func (s *AccountService) DeleteAccountByUserId(userId int64) error {
	return s.repo.DeleteAccountByUserId(userId)
}

func (s *AccountService) GetAccountByUserId(userId int64) (models.Account, error) {
	account,err:=s.repo.GetAccountByUserId(userId)
	if err!=nil{
		if errors.Is(err, sql.ErrNoRows) {
			accountID,err:=s.repo.CreateAccountByUserId(userId)
			if err!=nil{
				log.Printf("[ERROR]: %v",err)
				return account,err
			}
			log.Printf("[INFO]: account with id %d was created",accountID)
			return models.Account{Id: accountID,UserId: userId},nil
		}
	}
	return account,err
}
