package dblayer

import (
	"errors"
	"log"
)

const (
	mongo = "mongodb"
	msql  = "mysql"
)

var ErrDBTypeNotFound = errors.New("[ERROR] Database Type Not Found!")

type DBLayer interface {
	AddClub(ci *ClubInfo) error
	FindClub(id int) (ClubInfo, error)
	AllClubs() (club, error)
}

type ClubInfo struct {
	ID      int    `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Ranking string `json:"ranking" bson:"ranking(2017-2018)"`
	League  string `json:"league" bson:"league"`
}

type club []ClubInfo

// ConnectDatabase connects to a database type o using the provided connection string
func ConnectDatabase(o, c string) (DBLayer, error) {
	switch o {
	case mongo:
		return NewMongoDBDataStore(c)
	case msql:
		return NewMySQLDataStore(c)
	}
	log.Println("[INFO] Could not find", o)
	return nil, ErrDBTypeNotFound
}
