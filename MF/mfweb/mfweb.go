package mfweb

import (
	"-go-/_mfLogger"

	"fmt"
	"net/http"
)

func Run() {

	http.HandleFunc("/", sroot)

	// CASE1:
	http.Handle("/testhandle", newHandler())
	http.ListenAndServe(":8080", nil)

	// CASE2:
	// http.ListenAndServe(":8080", newHandler())
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := _mfLogger.GetInstance()
	fmt.Fprint(w, "Welcome to the MF Software System (sroot)")
	logger.Println("[INFO] Received an HTTP Get Request on Root URL")
}
