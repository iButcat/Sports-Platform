package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func WriteSportsDataToFile(data interface{}) error {
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
