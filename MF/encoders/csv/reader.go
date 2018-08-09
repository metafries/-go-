package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("cfile.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#'
	r.Comma = ';'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(records)

	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("    CSV Row", record)
	// }
}
