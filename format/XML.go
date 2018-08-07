package main

import (
	"encoding/xml"
	"fmt"
)

func main() {

	type CrewMember struct {
		ID                int      `xml:"id"`
		Name              string   `xml:"name,attr"`
		SecurityClearance int      `xml:"clearancelevel"`
		AccessCodes       []string `xml:"accesscodes>code"`
	}

	type ShipInfo struct {
		XMLName   xml.Name `xml:"SHIPINFO"`
		ShipID    int      `xml:"ShipDetails>ShipID"`
		ShipClass string   `xml:"ShipDetails>ShipClass"`
		Captain   CrewMember
	}

	cm3 := CrewMember{Name: "Cashviar", SecurityClearance: 20, AccessCodes: []string{"NEW", "TOT"}}
	si3 := ShipInfo{ShipID: 1, ShipClass: "Fighter", Captain: cm3}
	slices := []int{1, 2, 3, 4}

	bsi3, err := xml.MarshalIndent(&si3, "", "    ")
	PrintResult(err, bsi3)
	bslices, err := xml.Marshal(slices)
	PrintResult(err, bslices)

}

func PrintResult(err error, b []byte) {
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	fmt.Println(xml.Header, string(b))
}
