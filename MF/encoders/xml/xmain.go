package main

import (
	"encoding/xml"
	"log"
	"os"
)

type CrewMember struct {
	XMLName           xml.Name `xml:"member"`
	ID                int      `xml:"id,omitempty"`
	Name              string   `xml:"name,attr"`
	SecurityClearance int      `xml:"clearance,attr"`
	AccessCodes       []string `xml:"codes>code"`
}

type ShipInfo struct {
	XMLName   xml.Name `xml:"ship"`
	ShipID    int      `xml:"ShipInfo>ShipID"`
	ShipClass string   `xml:"ShipInfo>ShipClass"`
	Captain   CrewMember
}

func main() {

	file, err := os.Create("xfile.xml")
	if err != nil {
		log.Fatal("[Fatal] Could not create file -", err)
	}
	defer file.Close()

	cm3 := CrewMember{Name: "Cashviar", SecurityClearance: 20, AccessCodes: []string{"NEW", "TOT"}}
	si3 := ShipInfo{ShipID: 1, ShipClass: "Fighter", Captain: cm3}

	enc := xml.NewEncoder(file)
	enc.Indent(" ", "	")
	enc.Encode(si3)
	if err != nil {
		log.Fatal("[Fatal] Could not encode xml file -", err)
	}
}
