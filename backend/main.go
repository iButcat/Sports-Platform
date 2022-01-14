package main

import (
	"flag"
	"fmt"
	logSpecial "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// internal pkg

	"backend/api/service"
	"backend/api/transport"
	"backend/config"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/iButcat/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	var (
		httpAddr = flag.String("http.addr", ":8000", "HTTP listen address") // should  be 8080
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
		var err error
		db, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}
	}

	repository := repository.NewRepo(db, logSpecial.Logger{})
	/*
		url := config.URL
		sports, err := utils.FetchSportsAPI(url)
		if err != nil {
			logSpecial.Fatal(err)
		}
		ctx := context.Background()
		repository.Create(ctx, sports)
	*/

	var serviceImplt service.Service
	{
		serviceImplt = service.NewService(repository, logger)
	}

	var h http.Handler
	{
		h = transport.MakeHTTPHandler(serviceImplt, log.With(logger, "component", "HTTP"))
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
