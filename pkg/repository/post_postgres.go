package repository

import (
	"errors"
	"fmt"
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) CreatePost(post models.Post) (int64, error) {
	mu.Lock()
	defer mu.Unlock()

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var postId int64
	createPostQuery := fmt.Sprintf("INSERT INTO %s (title, description,created_at,account_id) values ($1, $2,$3, $4) RETURNING id", postTable)

	row := tx.QueryRow(createPostQuery, post.Title, post.Description, time.Now(), post.AccountId)
	err = row.Scan(&postId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	post.Id = postId
	createImagesQuery := fmt.Sprintf("INSERT INTO %s (link, post_id) VALUES ", imageTable)
	var inserts []string
	for _, image := range post.Images {
		inserts = append(inserts, fmt.Sprintf("('%s',%d)", image.Link, postId))
	}
	createImagesQuery += strings.Join(inserts, ",")

	_, err = tx.Exec(createImagesQuery)

	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error while adding images: %v", err)
	}

	return postId, tx.Commit()
}

func (r *PostPostgres) DeletePostById(postId int64) error {
	mu.Lock()
	defer mu.Unlock()
	var postExist bool

	err := r.db.QueryRowx(fmt.Sprintf("SELECT 1 FROM %s WHERE id=$1", postTable), postId).Scan(&postExist)
	if err != nil {
		return err
	}

	if !postExist {
		return errors.New("post doesn't exist")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteImagesQuery := fmt.Sprintf("DELETE FROM %s WHERE post_id=$1", imageTable)
	_, err = tx.Exec(deleteImagesQuery, postId)
	if err != nil {
		tx.Rollback()
		return err
	}

	deletePostQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postTable)
	_, err = tx.Exec(deletePostQuery, postId)

	return tx.Commit()
}