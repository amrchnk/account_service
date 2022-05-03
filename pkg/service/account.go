package service

import (
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/amrchnk/account_service/pkg/repository"
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
	return s.repo.GetAccountByUserId(userId)
}
