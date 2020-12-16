package api

import (
	"context"

	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

type repo struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) Create(ctx context.Context, post Post) (string, error) {
	repo.db.Debug().AutoMigrate(&post)
	if len(post.Title) == 0 || len(post.Article) == 0 || len(post.Image) == 0 {
		return "Cannot be empty", nil
	}
	if err := repo.db.Debug().Create(&post).Error; err != nil {
		return "", err
	}
	return "Saved", nil
}

func (repo *repo) Get(ctx context.Context, id string) (Post, error) {
	var post Post
	sqlGet := "SELECT id, title, article, image FROM posts WHERE id=?"
	if err := repo.db.Debug().Raw(sqlGet, id).Scan(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (repo *repo) GetAll(ctx context.Context) ([]Post, error) {
	var posts []Post
	sqlGetAll := "SELECT id, title, article, image FROM posts"
	if err := repo.db.Debug().Raw(sqlGetAll).Scan(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *repo) Update(ctx context.Context, post Post) (string, error) {
	sqlUpdate := "UPDATE posts SET title=?, article=?, image=? WHERE id=?"
	if err := repo.db.Exec(sqlUpdate, post.Title, post.Article, post.Image, post.ID).Error; err != nil {
		return "", err
	}
	return "Updated", nil
}

func (repo *repo) Delete(ctx context.Context, id string) (string, error) {
	if err := repo.db.Exec("DELETE FROM posts WHERE id=?", id).Error; err != nil {
		return "", err
	}
	return "Deleted", nil
}
