package mfportal

import (
	"-go-/MF/dblayer"
	"-go-/MF/mfweb/mfrestapi"
	"-go-/_mfConfig"
	"html/template"
	"log"
	"net/http"
)

var mfWebTemplate *template.Template

func Run() error {
	var err error
	mfWebTemplate, err = template.ParseFiles(
		"./mfweb/mfportal/cover/Club/club.html",
		"./mfweb/mfportal/cover/about/about.html")
	if err != nil {
		return err
	}
	conf := struct {
		Filespath string `json:"filespath"`
	}{}
	err = _mfConfig.GetConfiguration(_mfConfig.JSON, &conf, "./mfweb/portalconfig.json")
	if err != nil {
		return err
	}

	mfrestapi.IntializeAPIHandlers()
	log.Println(conf.Filespath)
	fs := http.FileServer(http.Dir(conf.Filespath))
	http.Handle("/", fs)
	http.HandleFunc("/Club/", clubhandler)
	http.HandleFunc("/about/", abouthandler)
	return http.ListenAndServe(":8061", nil)
}

func clubhandler(w http.ResponseWriter, r *http.Request) {
	dbl, err := dblayer.ConnectDatabase("mysql", "root:KLin#180812@/MF")
	if err != nil {
		return
	}
	all, err := dbl.AllClubs()
	if err != nil {
		return
	}
	err = mfWebTemplate.ExecuteTemplate(w, "club.html", all)
	if err != nil {
		log.Println(err)
	}
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	about := struct {
		Msg string `json:"message"`
	}{}
	err := _mfConfig.GetConfiguration(_mfConfig.JSON, &about, "./mfweb/about.json")
	if err != nil {
		return
	}
	err = mfWebTemplate.ExecuteTemplate(w, "about.html", about)
	if err != nil {
		log.Println(err)
	}
}
