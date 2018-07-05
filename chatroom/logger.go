package main

import (
	"-go-/_mfLogger"
	"fmt"
	"net/http"
)

func main() {
	logger := _mfLogger.GetInstance()
	logger.Println("Starting MetaFries web service")
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := _mfLogger.GetInstance()
	fmt.Fprint(w, "Welcome to the MetaFries software system.")
	logger.Println("Received an http GET request on root url")
}
