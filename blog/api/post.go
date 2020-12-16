package api

import "context"

type Post struct {
	ID      string `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Article string `json:"article"`
	Image   string `json:"image"`
}

type Repository interface {
	Create(ctx context.Context, post Post) (string, error)
	Get(ctx context.Context, id string) (Post, error)
	GetAll(ctx context.Context) ([]Post, error)
	Update(ctx context.Context, post Post) (string, error)
	Delete(ctx context.Context, id string) (string, error)
}
