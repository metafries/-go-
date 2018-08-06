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
	msk2 := map[string]int{"item1": 1, "item2": 2}
	mik2 := map[int]string{1: "item1", 2: "item2"} // non-string map keys (Go 1.7 -)
	scm := []CrewMember{cm4, cm3}

	b4, err := json.Marshal(&cm4)
	PrintResult(err, b4)
	b3, err := json.Marshal(&cm3)
	PrintResult(err, b3)
	bsi3, err := json.Marshal(&si3)
	PrintResult(err, bsi3)
	bmsk2, err := json.Marshal(&msk2)
	PrintResult(err, bmsk2)
	bmik2, err := json.Marshal(&mik2)
	PrintResult(err, bmik2)
	bscm, err := json.Marshal(&scm)
	PrintResult(err, bscm)

	sbyte := []byte(bsi3)
	mbyte := []byte(bmik2)
	dbyte := []byte(bscm)

	si := new(ShipInfo)
	m := make(map[int]string)
	s := []CrewMember{}

	json.Unmarshal(sbyte, si)
	json.Unmarshal(mbyte, &m)
	json.Unmarshal(dbyte, &s)

	fmt.Println(si.ShipID, si.ShipClass, si.Captain.Name)
	fmt.Println(m)
	fmt.Println(s)

}

func PrintResult(err error, b []byte) {
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}
	fmt.Println(string(b))
}
