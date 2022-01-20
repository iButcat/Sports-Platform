package service

import (
	"backend/models"
	"context"
	"fmt"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/iButcat/repository"
)

type Service interface {
	Get(ctx context.Context, id string) (models.Data, error)
	GetName(ctx context.Context, teams string) (models.Data, error)
	GetAll(ctx context.Context) (*models.Sports, error)
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

func (s service) Get(ctx context.Context, id string) (models.Data, error) {
	var sports models.Sports
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
	var data models.Data // use a slice instead if multiple result
	var datas []models.Data
	for _, value := range allData.(*models.Sports).Data {
		if value.Teams[0] == teams || strings.ToLower(value.Teams[1]) == teams {
			data = value
		} else if strings.HasPrefix(value.Teams[0], teams) || strings.HasPrefix(value.Teams[1], teams) {
			datas = append(datas, value)
		}
	}
	fmt.Println("multiple result: ", datas)
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
