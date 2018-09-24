package mfrestapi

import (
	"-go-/_mfConfig"
	"log"
	"net/http"
)

type DBlayerconfig struct {
	DB   string `json:"database"`
	Conn string `json:"connectionstring"`
}

func IntializeAPIHandlers() error {
	conf := new(DBlayerconfig)
	err := _mfConfig.GetConfiguration(_mfConfig.JSON, conf, "../../apiconfig.json")
	if err != nil {
		log.Println("ERROR: decoding JSON: ", err)
		return err
	}
	h := NewMFClubReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil {
		log.Println("ERROR: connecting to db: ", err)
		return err
	}
	http.HandleFunc("/mfclub/", h.handleMFClubReq)
	return nil
}

func RunAPI() error {
	if err := IntializeAPIHandlers(); err != nil {
		return err
	}
	return http.ListenAndServe(":8061", nil)
}
