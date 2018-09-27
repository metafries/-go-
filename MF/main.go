package main

import (
	"-go-/MF/mfweb/mfportal"
	"-go-/_mfLogger"
)

func main() {
	logger := _mfLogger.GetInstance()
	logger.Println("Starting MF Web Service ...")

	mfportal.Run()
}
