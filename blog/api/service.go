package api

import (
	"context"
)

type Service interface {
	CreatePost(ctx context.Context, title string, article string, image string) (string, error)
	GetPost(ctx context.Context, id string) (Post, error)
	GetAllPost(ctx context.Context) ([]Post, error)
	UpdatePost(ctx context.Context, post Post) (string, error)
	DeletePost(ctx context.Context, id string) (string, error)
}
