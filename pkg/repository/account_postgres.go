package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
	"time"
)

var mu = &sync.Mutex{}

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) CreateAccountByUserId(userId int64) (int64, error) {
	mu.Lock()
	defer mu.Unlock()
	var id int64
	CreateAccountQuery := fmt.Sprintf("INSERT INTO %s (user_id, created_at, profile_image) values ($1, $2, $3) RETURNING id", accountsTable)
	row := r.db.QueryRow(CreateAccountQuery, userId, time.Now(), defaultAvatarUrl)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AccountPostgres) DeleteAccountByUserId(userId int64) error {
	mu.Lock()
	defer mu.Unlock()
	DeleteAccountQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", accountsTable)
	_, err := r.db.Exec(DeleteAccountQuery, userId)

	return err
}

func (r *AccountPostgres) GetAccountByUserId(userId int64) (models.Account, error) {
	mu.Lock()
	defer mu.Unlock()

	var account models.Account
	GetAccountQuery := fmt.Sprintf("SELECT * FROM %s where user_id=$1", accountsTable)
	err := r.db.Get(&account, GetAccountQuery, userId)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("[ERROR]: %v",err)
		return account,errors.New("account doesn't exist")
	}
	return account, err
}

func (r *AccountPostgres) UpdateAccountInfo(updates models.UpdateAccountInfo) (int64, error) {
	mu.Lock()
	defer mu.Unlock()
	var accountId int64
	UpdateAccountQuery := fmt.Sprintf("UPDATE %s SET profile_image = $1 WHERE user_id = $2 RETURNING id", accountsTable)
	row := r.db.QueryRow(UpdateAccountQuery, updates.ProfileImage, updates.UserId)
	err := row.Scan(&accountId)

	if errors.Is(err, sql.ErrNoRows){
		log.Printf("[ERROR]: %v",err)
		return 0,errors.New("account doesn't exist")
	}

	if err != nil {
		return 0, err
	}
	return accountId, err
}
