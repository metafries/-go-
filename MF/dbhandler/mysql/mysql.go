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

	db, err := sql.Open("mysql", "root:KLin#180812@/MF?parseTime=true")
	if err != nil {
		log.Fatal("[FATAL] Could not connect, error -> ", err.Error())
	}
	defer db.Close()

	clubs := GetClubByLeagues(db, []string{"'Premier League'"})
	fmt.Println("    Club Table Data : ", clubs)

	fmt.Println("    Club ID = 1:      ", GetClubInfoById(db, 1))

	AddClubInfo(db, clubInfo{name: "Juventus", ranking: 1, league: "Lega Serie A"})

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

func GetClubInfoById(db *sql.DB, id int) (ci clubInfo) {

	row := db.QueryRow("SELECT * FROM Club WHERE id = ?", id)

	err := row.Scan(&ci.id, &ci.name, &ci.ranking, &ci.stadium, &ci.league)
	if err != nil {
		log.Fatal("[FATAL] -> ", err)
	}
	return
}

func AddClubInfo(db *sql.DB, ci clubInfo) int64 {

	resp, err := db.Exec(
		"INSERT INTO Club (name, ranking, league) VALUES (?, ?, ?)",
		ci.name, ci.ranking, ci.league)
	if err != nil {
		log.Fatal("[FATAL] -> ", err)
	}
	ra, _ := resp.RowsAffected()
	id, _ := resp.LastInsertId()

	log.Println("[INFO] -> Rows Affected:", ra, ", Last Insert Id:", id)
	return id

}
