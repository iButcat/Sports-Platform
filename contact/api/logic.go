package api

import (
	"context"

	"github.com/go-kit/kit/log"
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

func (s service) CreateContact(ctx context.Context, contact Contact) (string, error) {
	createRepo, err := s.repository.Create(ctx, contact)
	if err != nil {
		return "", err
	}
	return createRepo, nil
}

func (s service) GetContact(ctx context.Context, id string) (Contact, error) {
	getRepo, err := s.repository.Get(ctx, id)
	if err != nil {
		return getRepo, err
	}
	return getRepo, nil
}

func (s service) GetAllContact(ctx context.Context) ([]Contact, error) {
	getAllRepo, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return getAllRepo, nil
}

func (s service) DeleteContact(ctx context.Context, id string) (string, error) {
	deleteRepo, err := s.repository.Delete(ctx, id)
	if err != nil {
		return "", err
	}
	return deleteRepo, nil
}
