package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

/*
Our data structure which represent
json data from sports api.
*/

type Sports struct {
	gorm.Model
	Data []Data `gorm:"many2many:data_id" json:"data"`
}

type Data struct {
	gorm.Model
	ID           string         `gorm:"-" json:"id,omitempty"`
	SportKey     string         `json:"sport_key"`
	SportNice    string         `json:"sport_nice"`
	CommenceTime int            `json:"commence_time"`
	Teams        pq.StringArray `gorm:"type:varchar(64)[]" json:"teams"`
	HomeTeam     string         `json:"home_team"`
	Sites        []Sites        `gorm:"many2many:sites_id" json:"sites"`
}

type Sites struct {
	gorm.Model
	SiteKey    string `json:"site_key"`
	SiteNice   string `json:"site_nice"`
	LastUpdate int    `json:"last_update"`
	Odds       struct {
		/*
			Using external package for
			the correct type for
			saving it into the database
		*/
		H2H pq.Float64Array `gorm:"type:varchar(64)[]" json:"h2h"`
	} `gorm:"embedded" json:"odds"`
}
