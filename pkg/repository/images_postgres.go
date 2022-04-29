package repository

import (
	"fmt"
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type ImagesPostgres struct {
	db *sqlx.DB
}

func NewImagesPostgres(db *sqlx.DB) *ImagesPostgres {
	return &ImagesPostgres{db: db}
}

func (r *ImagesPostgres) GetImagesFromPost(postId int64) ([]models.Image, error) {
	var images []models.Image
	selectImagesQuery := fmt.Sprintf("SELECT * FROM %s WHERE post_id=$1", imageTable)
	err := r.db.Select(&images, selectImagesQuery, postId)

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return images, fmt.Errorf("error while getting images: %v", err)
	}
	return images, err
}
