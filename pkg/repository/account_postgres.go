package repository

import (
	"fmt"
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/jmoiron/sqlx"
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
	CreateAccountQuery := fmt.Sprintf("INSERT INTO %s (user_id,created_at) values ($1, $2) RETURNING id", accountsTable)
	row := r.db.QueryRow(CreateAccountQuery, userId, time.Now())
	if err := row.Scan(&id); err != nil {
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
	return account, err
}
