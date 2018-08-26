package dblayer

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type mySQLDataStore struct {
	*sql.DB
}

func NewMySQLDataStore(conn string) (*mySQLDataStore, error) {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	return &mySQLDataStore{DB: db}, nil
}

func (msql *mySQLDataStore) AddClub(ci *ClubInfo) error {
	_, err := msql.Exec(
		"INSERT INTO Club (name, ranking, league) VALUES (?, ?, ?)",
		ci.name, ci.ranking, ci.league)
	return err
}

func (msql *mySQLDataStore) FindClub(id int) (ClubInfo, error) {
	row := msql.QueryRow("SELECT * FROM Club WHERE id = ?", id)
	ci := ClubInfo{}
	err := row.Scan(&ci.id, &ci.name, &ci.ranking, &ci.stadium, &ci.league)
	return ci, err
}

func (msql *mySQLDataStore) AllClubs() (club, error) {
	rows, err := msql.Query("SELECT * FROM Club;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clubs := club{}
	for rows.Next() {
		club := ClubInfo{}
		err := row.Scan(&ci.id, &ci.name, &ci.ranking, &ci.stadium, &ci.league)
		if err == nil {
			clubs = append(clubs, club)
		}
	}

	err = rows.Err()
	return clubs, err
}
