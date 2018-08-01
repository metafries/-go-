package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	GenFileStatusRpt("testfile.txt")
}

func GenFileStatusRpt(fname string) {
	// Stat returns file info. It will return
	// an error if there is no file.
	filestats, err := os.Stat(fname)
	PrintFatalError(err)

	fmt.Println("    What's the file name?                        ", filestats.Name())
	fmt.Println("    Am I a directory?                            ", filestats.IsDir())
	fmt.Println("    What are the permissions?                    ", filestats.Mode())
	fmt.Println("    What's the file size?                        ", filestats.Size())
	fmt.Println("    When was the last time the file modified?    ", filestats.ModTime())
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file - ", err)
	}
}
