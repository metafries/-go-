package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("cfile.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comment = '#'
	// r.Comma = ';'

	// records, err := r.ReadAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(records)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("    CSV Row", record)

		i, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i * 4)
	}
}
