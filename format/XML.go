package main

import (
	"encoding/xml"
	"fmt"
)

func main() {

	type CrewMember struct {
		ID                int      `xml:"id,omitempty"`
		Name              string   `xml:"name"`
		SecurityClearance int      `xml:"clearancelevel"`
		AccessCodes       []string `xml:"accesscodes"`
	}

	type ShipInfo struct {
		ShipID    int
		ShipClass string
		Captain   CrewMember
	}

	cm3 := CrewMember{Name: "Cashviar", SecurityClearance: 20, AccessCodes: []string{"NEW", "TOT"}}
	si3 := ShipInfo{1, "Fighter", cm3}

	bsi3, err := xml.MarshalIndent(&si3, "", "    ")
	PrintResult(err, bsi3)

}

func PrintResult(err error, b []byte) {
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	fmt.Println(string(b))
}
