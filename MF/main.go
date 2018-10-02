package main

import (
	"-go-/MF/mfweb/mfportal"
	"-go-/_mfChat"
	"-go-/_mfLogger"
	"flag"
	"strings"
)

func main() {
	logger := _mfLogger.GetInstance()
	logger.Println("Starting MF Web Service ...")
	operation := flag.String("o", "w", "Operation: w for web \n c for chat")
	flag.Parse()
	switch strings.ToLower(*operation) {
	case "c":
		err := _mfChat.Run(":2100")
		if err != nil {
			logger.Println("ERROR: Could not run MF chat", err)
		}
	case "w":
		err := mfportal.Run()
		if err != nil {
			logger.Println("ERROR: Could not run MF web portal", err)
		}
	}
}
