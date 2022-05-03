package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/amrchnk/account_service/pkg/models"
	"github.com/jmoiron/sqlx"
	"log"
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
		log.Printf("[ERROR]: %v", err)
		tx.Rollback()
		return 0, err
	}

	var postId int64
	createPostQuery := fmt.Sprintf("INSERT INTO %s (title, description,created_at, updated_at, account_id) values ($1, $2,$3, $4, $5) RETURNING id", postTable)

	row := tx.QueryRow(createPostQuery, post.Title, post.Description, time.Now(), time.Now(), post.AccountId)
	err = row.Scan(&postId)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
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
		log.Printf("[ERROR]: %v", err)
		tx.Rollback()
		return 0, fmt.Errorf("error while adding images: %v", err)
	}

	if len(post.Categories) == 0 {
		createPostCategoriesQuery := fmt.Sprintf("INSERT INTO %s (category_id, post_id) VALUES ", postsCategoriesTable)
		createPostCategoriesQuery += fmt.Sprintf("('%d',%d)", 1, postId)

		_, err = tx.Exec(createPostCategoriesQuery)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			tx.Rollback()
			return 0, fmt.Errorf("error while adding categories to post: %v", err)
		}
	} else {
		createPostCategoriesQuery := fmt.Sprintf("INSERT INTO %s (category_id, post_id) VALUES ", postsCategoriesTable)
		var inserts []string
		for _, category := range post.Categories {
			inserts = append(inserts, fmt.Sprintf("('%d',%d)", category, postId))
		}

		createPostCategoriesQuery += strings.Join(inserts, ",")

		_, err = tx.Exec(createPostCategoriesQuery)

		if err != nil {
			log.Printf("[ERROR]: %v", err)
			tx.Rollback()
			return 0, fmt.Errorf("error while adding categories to post: %v", err)
		}
	}

	return postId, tx.Commit()
}

func (r *PostPostgres) DeletePostById(postId int64) error {
	mu.Lock()
	defer mu.Unlock()
	var postExist bool

	err := r.db.QueryRowx(fmt.Sprintf("SELECT 1 FROM %s WHERE id=$1", postTable), postId).Scan(&postExist)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("[ERROR]: %v", err)
			return fmt.Errorf("post doesn't exist")
		}

		log.Printf("[ERROR]: %v", err)
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteImagesQuery := fmt.Sprintf("DELETE FROM %s WHERE post_id=$1", imageTable)
	_, err = tx.Exec(deleteImagesQuery, postId)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		tx.Rollback()
		return err
	}

	deletePostQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postTable)
	_, err = tx.Exec(deletePostQuery, postId)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PostPostgres) GetPostById(postId int64) (models.Post, error) {
	mu.Lock()
	defer mu.Unlock()
	var post models.Post
	var images []models.Image
	var postExist bool

	err := r.db.QueryRowx(fmt.Sprintf("SELECT 1 FROM %s WHERE id=$1", postTable), postId).Scan(&postExist)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("[ERROR]: %v", err)
			return post, fmt.Errorf("post doesn't exist")
		}

		log.Printf("[ERROR]: %v", err)
		return post, err
	}

	selectImagesQuery := fmt.Sprintf("SELECT * FROM %s WHERE post_id=$1", imageTable)
	err = r.db.Select(&images, selectImagesQuery, postId)

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return post, err
	}

	var categories []uint8
	selectCategoriesQuery := fmt.Sprintf("SELECT category_id FROM %s WHERE post_id=$1", postsCategoriesTable)
	err = r.db.Select(&categories, selectCategoriesQuery, postId)

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return post, err
	}

	if len(categories) > 0 {
		for _, category := range categories {
			post.Categories = append(post.Categories, int64(category))
		}
	}

	selectPostQuery := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", postTable)
	err = r.db.Get(&post, selectPostQuery, postId)

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return post, err
	}

	post.Images = images

	return post, nil
}

func (r *PostPostgres) UpdatePostByd(post models.Post) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	var postExist bool
	err := r.db.QueryRowx(fmt.Sprintf("SELECT 1 FROM %s WHERE id=$1", postTable), post.Id).Scan(&postExist)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("[ERROR]: %v", err)
		return "", errors.New("post doesn't exist")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	if len(post.Images) != 0 {
		deleteImagesQuery := fmt.Sprintf("DELETE FROM %s WHERE post_id=$1", imageTable)
		_, err = tx.Exec(deleteImagesQuery, post.Id)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			tx.Rollback()
			return "", err
		}

		createImagesQuery := fmt.Sprintf("INSERT INTO %s (link, post_id) VALUES ", imageTable)
		var inserts []string
		for _, image := range post.Images {
			inserts = append(inserts, fmt.Sprintf("('%s',%d)", image.Link, post.Id))
		}
		createImagesQuery += strings.Join(inserts, ",")

		_, err = tx.Exec(createImagesQuery)

		if err != nil {
			log.Printf("[ERROR]: %v", err)
			tx.Rollback()
			return "", fmt.Errorf("error while updating images: %v", err)
		}
	}

	if len(post.Categories) != 0 {
		deletePostsCategoriesQuery := fmt.Sprintf("DELETE FROM %s WHERE post_id=$1", postsCategoriesTable)
		_, err = tx.Exec(deletePostsCategoriesQuery, post.Id)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			tx.Rollback()
			return "", err
		}

		createPostsCategoriesQuery := fmt.Sprintf("INSERT INTO %s (category_id, post_id) VALUES ", postsCategoriesTable)
		var inserts []string
		for _, category := range post.Categories {
			inserts = append(inserts, fmt.Sprintf("('%d',%d)", category, post.Id))
		}
		createPostsCategoriesQuery += strings.Join(inserts, ",")

		_, err = tx.Exec(createPostsCategoriesQuery)

		if err != nil {
			log.Printf("[ERROR]: %v", err)
			tx.Rollback()
			return "", fmt.Errorf("error while updating categories: %v", err)
		}
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	setValues = append(setValues, fmt.Sprintf("updated_at=$%d", argId))
	post.UpdatedAt = time.Now()
	args = append(args, post.UpdatedAt)
	argId++

	if post.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, post.Title)
		argId++
	}

	if post.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, post.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s p SET %s WHERE p.id = %d`,
		postTable, setQuery, post.Id)

	_, err = r.db.Exec(query, args...)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return "", err
	}

	return fmt.Sprintf("Post with id = %d was updated", post.Id), tx.Commit()
}

func (r *PostPostgres) GetPostsByUserId(userId int64) ([]models.Post, error) {
	mu.Lock()
	defer mu.Unlock()
	var posts []models.Post
	var accountId int64

	err := r.db.QueryRowx(fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1", accountsTable), userId).Scan(&accountId)
	if errors.Is(err, sql.ErrNoRows) {
		var id int64
		CreateAccountQuery := fmt.Sprintf("INSERT INTO %s (user_id,created_at) values ($1, $2) RETURNING id", accountsTable)
		row := r.db.QueryRow(CreateAccountQuery, userId, time.Now())
		if err = row.Scan(&id); err != nil {
			return posts, err
		}
		return posts, nil
	}

	if err != nil {
		return posts, err
	}

	selectPostsQuery := fmt.Sprintf("SELECT id, title, description, created_at, account_id FROM %s WHERE account_id=$1", postTable)
	err = r.db.Select(&posts, selectPostsQuery, accountId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("[ERROR]: %v", err)
			return posts, errors.New("account doesn't exist")
		}
		return posts, err
	}

	for index := range posts {
		var images []models.Image
		selectImagesQuery := fmt.Sprintf("SELECT * FROM %s WHERE post_id=$1", imageTable)
		err = r.db.Select(&images, selectImagesQuery, posts[index].Id)

		if err != nil {
			return posts, err
		}
		posts[index].Images = images

		categories := make([]uint8, 0, len(posts[index].Categories))
		selectCategoriesQuery := fmt.Sprintf("SELECT category_id FROM %s WHERE post_id=$1", postsCategoriesTable)
		err = r.db.Select(&categories, selectCategoriesQuery, posts[index].Id)

		if err != nil {
			log.Printf("[ERROR]: %v", err)
			return posts, err
		}

		if len(categories) > 0 {
			for _, category := range categories {
				posts[index].Categories = append(posts[index].Categories, int64(category))
			}
		}
	}
	return posts, err
}

func (r *PostPostgres) GetAllUsersPosts(offset, limit int64, sorting string) ([]models.GetAllUsersPosts, error) {
	mu.Lock()
	defer mu.Unlock()

	var posts []models.GetAllUsersPosts

	selectPostsQuery := fmt.Sprintf("SELECT p.id as id, p.title as title, p.description as description, p.created_at as created_at, ac.user_id as user_id FROM %s p INNER JOIN %s ac ON p.account_id=ac.id", postTable, accountsTable)
	switch sorting {
	case "asc":
		selectPostsQuery += fmt.Sprintf(" ORDER BY p.created_at")
	default:
		selectPostsQuery += fmt.Sprintf(" ORDER BY p.created_at DESC")
	}

	selectPostsQuery += fmt.Sprintf(" OFFSET %d LIMIT %d", offset, limit)

	err := r.db.Select(&posts, selectPostsQuery)
	if err != nil {
		log.Printf("[ERROR]: %v", err)
		return nil, err
	}

	for index := range posts {
		selectImagesQuery := fmt.Sprintf("SELECT link as image FROM %s where post_id=$1", imageTable)
		err := r.db.Select(&posts[index].Images, selectImagesQuery, posts[index].Id)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			return nil, err
		}

		selectCategoriesQuery := fmt.Sprintf("SELECT c.title as category FROM category c INNER JOIN posts_have_categories phc ON c.id=phc.category_id INNER JOIN post p ON p.id=phc.post_id WHERE p.id=$1")
		err = r.db.Select(&posts[index].Categories, selectCategoriesQuery, posts[index].Id)
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			return nil, err
		}
	}

	return posts, nil
}
