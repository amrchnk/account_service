package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	accountsTable = "account"
	postTable     = "post"
	imageTable    = "image"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	params := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", params)
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}

	return db, nil
}
