package main

import (
  _"fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "github.com/jinzhu/gorm"
  "github.com/lib/pq"
)

/*
// Struct data structure to receive Json data
type Sports struct {
  gorm.Model
  Success bool `json:"success"`
	Data    []struct {
		SportKey     string   `json:"sport_key"`
		SportNice    string   `json:"sport_nice"`
    Teams pq.StringArray `gorm:"type:varchar(64)[]" json:"teams"`
		CommenceTime int      `json:"commence_time"`
		HomeTeam     string   `json:"home_team"`
		Sites        []struct {
			SiteKey    string `json:"site_key"`
			SiteNice   string `json:"site_nice"`
			LastUpdate int    `json:"last_update"`
			Odds       struct {
				H2H pq.Float64Array `gorm:"type:varchar(64)[]" json:"h2h"`
			} `gorm:"embedded" json:"odds"`
		} `gorm:"embedded" json:"sites"`
	} `gorm:"embedded" json:"data"`
}
*/

type Sports struct {
  gorm.Model
  Data []Data `gorm:"many2many:data_id" json:"data"`
}

type Data struct {
  gorm.Model
  SportKey     string   `json:"sport_key"`
  SportNice    string   `json:"sport_nice"`
  CommenceTime int      `json:"commence_time"`
  Teams pq.StringArray `gorm:"type:varchar(64)[]" json:"teams"`
  HomeTeam     string   `json:"home_team"`
  Sites []Sites `gorm:"many2many:sites_id" json:"sites"`
}

type Sites struct {
  gorm.Model
  SiteKey    string `json:"site_key"`
  SiteNice   string `json:"site_nice"`
  LastUpdate int    `json:"last_update"`
  Odds       struct {
    H2H pq.Float64Array `gorm:"type:varchar(64)[]" json:"h2h"`
  } `gorm:"embedded" json:"odds"`
}

func getAllSports(url string, sports interface{}) {
  resp, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  responseData, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  json.Unmarshal(responseData, &sports)
}
