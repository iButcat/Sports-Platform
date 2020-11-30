package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"backend/api" // internal pkg
	"backend/config"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	config, err := config.LoadConfig("./config")
	if err != nil {
		fmt.Println("can't load config:", err)
	}

	var db *gorm.DB
	{
		// Our database credentials
		var err error
		db, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}
	}

	var service api.Service
	{
		repository := api.NewRepo(db, logger)
		service = api.NewService(repository, logger)
	}

	var h http.Handler
	{
		h = api.MakeHTTPHandler(service, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)

}
