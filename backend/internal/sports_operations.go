package internal

import (
	"backend/config"
	"backend/models"
	"context"
	"log"
	"strconv"

	"github.com/iButcat/repository"
)

type Operation struct {
	config     config.Config
	repository repository.Repository
}

func NewOperation(config config.Config, repository repository.Repository) *Operation {
	return &Operation{
		config:     config,
		repository: repository,
	}
}

func (op *Operation) CreateSports(sports models.Sports) error {
	ctx := context.Background()
	if _, err := op.repository.Create(ctx, &sports); err != nil {
		return err
	}
	return nil
}

func (op *Operation) UpdateSports() error {
	fetchSports, err := fetchSportsAPI(op.config.URL)
	if err != nil {
		return err
	}
	ctx := context.Background()
	var fields = make(map[string]interface{})
	for _, data := range fetchSports.Data {
		for _, site := range data.Sites {
			fields["h2_h"] = site.Odds.H2H
			var s string = strconv.FormatUint(uint64(site.ID), 10)
			if _, err := op.repository.Update(ctx, models.Sites{}, s, fields); err != nil {
				return err
			}
		}
	}
	return nil
}

func (op *Operation) DeleteSports() error {
	ctx := context.Background()
	allSports, err := op.repository.GetAll(ctx, &models.Sports{})
	if err != nil {
		log.Println(err)
	}
	sports := allSports.(*models.Sports)

	for _, data := range sports.Data {
		// TODO: Format CommenceTime to check if the match is finished and delete
		if data.CommenceTime > 105 {
			_, err := op.repository.Delete(ctx, &models.Data{}, data.ID)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}
