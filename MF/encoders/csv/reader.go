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
	r.Comma = ';'

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
			if pe, ok := err.(*csv.ParseError); ok {
				fmt.Println("\n    Bad Line:          ", pe.Line)
				fmt.Println("    Bad Column:        ", pe.Column)
				fmt.Println("    Error Reported:    ", pe.Err, "\n")
				if pe.Err == csv.ErrFieldCount {
					continue
				}
			}
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
