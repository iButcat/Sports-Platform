package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	// internal pkg
	"backend/utils"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repository Repository
	logger     log.Logger
}

// return new service
func NewService(repo Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

var sports Sports

// fetch sports from the api and save it to the database
func (s service) FetchSportsAPI(ctx context.Context, url string) (string, error) {
	logger := log.With(s.logger, "method", "FetchSports")
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(responseData, &sports); err != nil {
		return "", err
	}
	saveSports, err := s.repository.SaveSportsFetch(ctx, sports)
	if err != nil {
		return "", err
	}
	logger.Log(saveSports)
	if err := utils.WriteSportsDataToFile(&sports); err != nil {
		return "", err
	}
	return "Saved", nil
}

func (s service) Get(ctx context.Context, id string) (Data, error) {
	logger := log.With(s.logger, "method", "GetSports")
	data, err := s.repository.Get(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return data, err
	}
	return data, nil
}

func (s service) GetAll(ctx context.Context) ([]Data, []Sites, error) {
	logger := log.With(s.logger, "method", "GetAllSports")
	allData, allSites, err := s.repository.GetAll(ctx)
	if err != nil {
		level.Error(logger).Log("%s", err)
		return nil, nil, err
	}
	return allData, allSites, nil
}

func (s service) Update(ctx context.Context, sites Sites) (string, error) {
	logger := log.With(s.logger, "method", "UpdateSports")
	msg, err := s.repository.Update(ctx, sites)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return msg, nil
}

func (s service) Delete(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteSports")
	msg, err := s.repository.Delete(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return msg, nil
}
