package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8000/1ei30n11"
	resp, err := http.Get(url)
	inspectResponse(resp, err)

	data, err := json.Marshal(struct {
		X int
		Y float32
	}{X: 4, Y: 3.8})
	if err != nil {
		log.Fatal("FATAL: Error Occured While Marshaling JSON ", err)
	}
	resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	inspectResponse(resp, err)
}

func inspectResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatal("FATAL: Error Occured While Marshaling JSON ", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("FATAL: Error Occured While Trying to Read HTTP Response Body ", err)
	}
	log.Println(string(b))
}
