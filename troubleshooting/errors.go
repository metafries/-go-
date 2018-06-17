package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errCrewNotFound = errors.New("Crew member not found")

var scMapping = map[string]int{
	"RUS-KSA": 5,
	"EGY-URU": 1,
	"MRC-IRN": 1,
	"POR-ESP": 6,
	"FRA-AUS": 3,
	"ARG-ISL": 2,
	"PER-DNK": 1,
	"CRO-NGA": 2,
}

type findError struct {
	Game, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findSC(game, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(15)) * time.Second) // Simulate seraching
	if v, ok := scMapping[game]; !ok {
		return -1, errors.New("Crew member not found")
	} else {
		return v, nil
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	clearance, err := findSC("CRC-SRB", "server 1")
	fmt.Println("Clearance level found: ", clearance)
	fmt.Println("Error code: ", err)
}
