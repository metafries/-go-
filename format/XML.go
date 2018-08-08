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

	data := []byte(` 
	<SHIPINFO>
		<ShipDetails>
			<ShipID>1</ShipID>
			<ShipClass>Fighter</ShipClass>
		</ShipDetails>
		<Captain name="Cashviar">
			<id>0</id>
			<clearancelevel>20</clearancelevel>
			<accesscodes>
				<code>NEW</code>
				<code>TOT</code>
			</accesscodes>
		</Captain>
	</SHIPINFO>`)
	si := ShipInfo{}
	xml.Unmarshal(data, &si)
	fmt.Println(si)

}

func PrintResult(err error, b []byte) {
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	fmt.Println(xml.Header, string(b))
}
