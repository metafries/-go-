package main

import (
	"encoding/json"
	"fmt"
)

func main() {

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

	cm4 := CrewMember{1, "MetaFries", 10, []string{"GIR", "TOT"}}
	cm3 := CrewMember{Name: "Cashviar", SecurityClearance: 20, AccessCodes: []string{"NEW", "TOT"}}
	si3 := ShipInfo{1, "Fighter", cm3}

	b4, err := json.Marshal(&cm4)
	PrintError(err)
	b3, err := json.Marshal(&cm3)
	PrintError(err)
	bsi3, err := json.Marshal(&si3)
	PrintError(err)

	fmt.Println(string(b4))
	fmt.Println(string(b3))
	fmt.Println(string(bsi3))

}

func PrintError(err error) {
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
}
