package api

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

func (s service) CreatePost(ctx context.Context, title string, article string, image string) (string, error) {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	post := Post{
		ID:      id,
		Title:   title,
		Article: article,
		Image:   image,
	}
	fmt.Println(post)
	savePost, err := s.repository.Create(ctx, post)
	if err != nil {
		return "", err
	}
	return savePost, nil
}

func (s service) GetPost(ctx context.Context, id string) (Post, error) {
	getPost, err := s.repository.Get(ctx, id)
	if err != nil {
		return getPost, err
	}
	return getPost, nil
}

func (s service) GetAllPost(ctx context.Context) ([]Post, error) {
	getAllPosts, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return getAllPosts, nil
}

func (s service) UpdatePost(ctx context.Context, post Post) (string, error) {
	post = Post{
		ID:      post.ID,
		Title:   post.Title,
		Article: post.Article,
		Image:   post.Image,
	}
	updatePost, err := s.repository.Update(ctx, post)
	if err != nil {
		return "", err
	}
	return updatePost, nil
}

func (s service) DeletePost(ctx context.Context, id string) (string, error) {
	deletePost, err := s.repository.Delete(ctx, id)
	if err != nil {
		return "", err
	}
	return deletePost, nil
}
