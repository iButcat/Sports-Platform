package utils

import (
	"backend/api/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func writeSportsDataToFile(data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	if err := createFolderIfNotExists("./data"); err != nil {
		return err
	}

	if err := ioutil.WriteFile("data/sports.json", file, 0644); err != nil {
		return err
	}
	return nil
}

func createFolderIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

// fetch sports from the api and save it into the database
func FetchSportsAPI(url string) (*models.Sports, error) {

	var sports models.Sports

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(responseData, &sports); err != nil {
		return nil, err
	}

	if err := writeSportsDataToFile(&sports); err != nil {
		return nil, err
	}
	return &sports, nil
}
