package test

import (
	"context"
	"fmt"
	"os"

	"blog/api" // internal pkg
	"blog/config"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setup our service
func setup() (service api.Service, ctx context.Context) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var db *gorm.DB
	{
		var err error
		config, err := config.LoadConfig("../config")
		if err != nil {
			fmt.Println("can't load config:", err)
		}

		db, err = gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}
	}
	repository := api.NewRepo(db, logger)
	service = api.NewService(repository, logger)
	return service, context.Background()
}
