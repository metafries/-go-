package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type clubInfo struct {
	id      int
	name    string
	ranking int
	stadium string
	league  string
}

type Club []clubInfo

func main() {

	db, err := sql.Open("mysql", "root:<password>@/MF?parseTime=true")
	if err != nil {
		log.Fatal("[FATAL] Could not connect, error -> ", err.Error())
	}
	defer db.Close()

	clubs := GetClubByLeagues(db, []string{"'Premier League'"})
	fmt.Println("    Club Table Data : ", clubs)
}

func GetClubByLeagues(db *sql.DB, leagues []string) Club {

	q := fmt.Sprintf("SELECT * FROM Club WHERE league in (%s);", strings.Join(leagues, ","))

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal("[FATAL] Could not get data from the Club table -> ", err)
	}
	defer rows.Close()

	retVal := Club{}
	cols, _ := rows.Columns()
	fmt.Println("    Columns Detected :", cols)

	for rows.Next() {
		club := clubInfo{}
		err = rows.Scan(&club.id, &club.name, &club.ranking, &club.stadium, &club.league)
		if err != nil {
			log.Fatal("[FATAL] Error scanning row -> ", err)
		}
		retVal = append(retVal, club)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("[FATAL] -> ", err)
	}
	return retVal
}
