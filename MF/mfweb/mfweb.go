package mfweb

import (
	"-go-/_mfLogger"
	"time"

	"fmt"
	"net/http"
)

func Run() {

	http.HandleFunc("/", sroot)

	// CASE1:
	// http.Handle("/testhandle", newHandler())
	// http.ListenAndServe(":8080", nil)

	// CASE2:
	// http.ListenAndServe(":8080", newHandler())

	// CASE3:
	http.HandleFunc("/testquery", queryTestHandler)
	http.Handle("/testhandle", newHandler())
	http.ListenAndServe(":8080", nil)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      newHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	server.ListenAndServe()
}

func queryTestHandler(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	message := fmt.Sprintf("Query map: %v \n", q)

	v1, v2 := q.Get("key1"), q.Get("key2")
	if v1 == v2 {
		message = message + fmt.Sprintf("v1 & v2 are equal to %s \n", v1)
	} else {
		message = message + fmt.Sprintf("v1 is equal to %s, v2 is equal to %s \n", v1, v2)
	}
	fmt.Fprint(w, message)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := _mfLogger.GetInstance()
	fmt.Fprint(w, "Welcome to the MF Software System (sroot)")
	logger.Println("[INFO] Received an HTTP Get Request on Root URL")
}
