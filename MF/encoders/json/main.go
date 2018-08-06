package main

import (
	"encoding/json"
	"log"
	"os"
)

type CrewMember struct {
	ID                int      `json:"id,omitempty"`
	Name              string   `json:"name"`
	SecurityClearance int      `json:"clearancelevel"`
	AccessCodes       []string `json:"accesscodes"`
}

type ShipInfo struct {
	ShipID    int
	ShipClass string
	Captain   CrewMember
}

func main() {
	f, err := os.Create("jfile.json")
	PrintFatalError(err)
	defer f.Close()

	// create some data to encode
	cm3 := CrewMember{Name: "Cashviar", SecurityClearance: 20, AccessCodes: []string{"NEW", "TOT"}}
	si3 := ShipInfo{1, "Fighter", cm3}

	err = json.NewEncoder(f).Encode(&si3)
	PrintFatalError(err)
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file - ", err)
	}
}
