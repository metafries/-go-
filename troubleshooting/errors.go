package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// ErrCrewNotFound <-Share between packages
var ErrCrewNotFound = errors.New("Crew member not found")

var scMapping = map[string]int{
	"RUS-KSA":  5,
	"EGY-URU":  1,
	"*MRC-IRN": 1,
	"POR-ESP":  6,
	"FRA-AUS":  3,
	"ARG-ISL":  2,
	"PER-DNK":  1,
	"CRO-NGA":  2,
	"CRC-SRB":  1,
	"GER-MEX":  1,
	"*BRA-SUI": 2,
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
		return -1, fmt.Errorf("%s could not be found on server '%s'", game, server)
	} else {
		return v, nil
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if clearance, err := findSC("BRA-SUI", "server 1"); err != nil {
		fmt.Println("Error Occured: [Type]", ErrCrewNotFound, "[Msg]", err)
	} else {
		fmt.Println("Clearance level found: ", clearance)
	}
}
