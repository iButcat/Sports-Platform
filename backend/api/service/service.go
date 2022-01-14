package service

import (
	"backend/api/models"
	"context"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/iButcat/repository"
)

type Service interface {
	Get(ctx context.Context, id string) (models.Data, error)
	GetName(ctx context.Context, teams string) (models.Data, error)
	GetAll(ctx context.Context) (*models.Sports, error)
	Update(ctx context.Context, sports models.Sports) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type service struct {
	repository repository.Repository
	logger     log.Logger
}

// return new service
func NewService(repo repository.Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

var sports models.Sports

func (s service) Get(ctx context.Context, id string) (models.Data, error) {
	logger := log.With(s.logger, "method", "GetSports")
	data, err := s.repository.Get(ctx, &sports, map[string]interface{}{"id": id})
	if err != nil {
		level.Error(logger).Log("err", err)
		return data.(models.Data), err
	}
	return data.(models.Data), nil
}

func (s service) GetName(ctx context.Context, teams string) (models.Data, error) {
	logger := log.With(s.logger, "method", "GetByName")
	allData, err := s.repository.GetAll(ctx, &models.Sports{})
	if err != nil {
		level.Error(logger).Log("err", err)
		return allData.(models.Data), err
	}
	var data models.Data
	for _, value := range allData.(*models.Sports).Data {
		if value.Teams[0] == teams || strings.ToLower(value.Teams[1]) == teams {
			data = value
			break
		}
	}
	return data, nil
}

func (s service) GetAll(ctx context.Context) (*models.Sports, error) {
	logger := log.With(s.logger, "method", "GetAllSports")
	sports, err := s.repository.GetAll(ctx, &models.Sports{})
	if err != nil {
		level.Error(logger).Log("%s", err)
		return nil, err
	}
	return sports.(*models.Sports), nil
}

func (s service) Update(ctx context.Context, sports models.Sports) (bool, error) {
	logger := log.With(s.logger, "method", "UpdateSports")
	updated, err := s.repository.Update(ctx, models.Sports{}, "1", map[string]interface{}{})
	if err != nil {
		level.Error(logger).Log("err", err)
		return updated, err
	}
	return updated, nil
}

func (s service) Delete(ctx context.Context, id string) (bool, error) {
	logger := log.With(s.logger, "method", "DeleteSports")
	// TODO: change message for bool instead
	_, err := s.repository.Delete(ctx, models.Sports{}, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return false, err
	}
	return true, nil
}
