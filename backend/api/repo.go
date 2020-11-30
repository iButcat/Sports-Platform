package api

import (
	"context"
	"fmt"
	"time"

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

func (repo *repo) SaveSportsFetch(ctx context.Context, sports Sports) (string, error) {
	var data Data
	var sites Sites
	sqlExec := "DROP TABLE if exists sports, data, sites, sites_id, data_id CASCADE"
	if err := repo.db.Exec(sqlExec).Error; err != nil {
		return "", err
	}
	repo.db.AutoMigrate(&sports, &data, &sites)
	if err := repo.db.Save(&sports).Error; err != nil {
		return "", err
	}
	return "Saved", nil
}

func (repo *repo) Get(ctx context.Context, id string) (Data, error) {
	var data Data
	var sites Sites
	sqlRaw := "SELECT * FROM data JOIN sites_id ON data.id = sites_id.data_id JOIN sites ON data.id = sites.id"
	if err := repo.db.Raw(sqlRaw).Scan(&data).Scan(&sites).Error; err != nil {
		return data, err
	}
	fmt.Println(data)
	return data, nil
}

func (repo *repo) GetAll(ctx context.Context) ([]Data, []Sites, error) {
	sqlRaw := "SELECT * FROM data JOIN sites_id ON data.id = sites_id.data_id JOIN sites ON sites_id.data_id = sites.id"
	var data []Data
	if err := repo.db.Debug().Raw(sqlRaw).Scan(&data).Error; err != nil {
		return nil, nil, err
	}
	var sites []Sites
	if err := repo.db.Debug().Raw(sqlRaw).Scan(&sites).Error; err != nil {
		return nil, nil, err
	}
	return data, sites, nil
}

// This function should update Odds every hours
func (repo *repo) Update(ctx context.Context, sites Sites) (string, error) {
	if err := repo.db.Exec("UPDATE sites SET h2_h=?", sites.Odds).Error; err != nil {
		return "", err
	}
	return "Success", nil
}

func (repo *repo) Delete(ctx context.Context, id string) (string, error) {
	sqlExec := "DROP TABLE if exists sports, data, sites, sites_id, data_id CASCADE"
	now := time.Now()
	//tomorrow := now.Add(24 * time.Hour)
	checkDate := "SELECT * FROM data WHERE CAST(created_at as time) != ?;"
	if err := repo.db.Debug().Exec(checkDate, now).Error; err != nil {
		repo.db.Exec(sqlExec)
		return "", err
	}
	/*
		else {
			if err := repo.db.Exec(sqlExec).Error; err != nil {
				return "", err
			}

		}*/
	return "Sucess", nil
}
