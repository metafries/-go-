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
		// panic("Crew member not found")
		return -1, findError{
			game,
			server,
			"Crew member not found",
		}
	} else {
		return v, nil
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	/* 	defer func() {
	   		if err := recover(); err != nil {
	   			fmt.Println("A panic recovered", err)
	   		}
	   	}()
	*/if clearance, err := findSC("BRA-SUI", "server 1"); err != nil {
		if v, ok := err.(findError); ok {
			fmt.Println(
				"Error Occured: [Type]", ErrCrewNotFound,
				"[Msg]", v.Game, "could not be found on", v.Server)
		}
	} else {
		fmt.Println("Clearance level found: ", clearance)
	}
}
