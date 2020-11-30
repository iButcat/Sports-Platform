package test

import (
	"fmt"
	"testing"

	// internal pkg
	"backend/config"

	"github.com/stretchr/testify/assert"
)

func TestFetchAPIService(t *testing.T) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		fmt.Println("can't load config:", err)
	}
	service, ctx := setup()
	fetchAPI, err := service.FetchSportsAPI(ctx, config.URL)
	if err != nil {
		return
	}
	assert.Equal(t, fetchAPI, "Saved")
}

func TestGetService(t *testing.T) {

}

func TestGetAllService(t *testing.T) {

}
