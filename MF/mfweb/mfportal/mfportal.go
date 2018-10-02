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

	conf := struct {
		Filespath string   `json:"filespath"`
		Templates []string `json:"templates"`
	}{}
	err = _mfConfig.GetConfiguration(_mfConfig.JSON, &conf, "./mfweb/portalconfig.json")
	if err != nil {
		return err
	}

	mfWebTemplate, err = template.ParseFiles(conf.Templates...)
	if err != nil {
		log.Println(err)
		return err
	}

	mfrestapi.IntializeAPIHandlers()
	log.Println(conf.Filespath)
	fs := http.FileServer(http.Dir(conf.Filespath))
	http.Handle("/", fs)
	http.HandleFunc("/Club/", clubhandler)
	http.HandleFunc("/about/", abouthandler)
	http.HandleFunc("/chat/", chathandler)
	return http.ListenAndServe(":8061", nil)
}

func chathandler(w http.ResponseWriter, r *http.Request) {
	ns := struct {
		Name string
	}{}
	r.ParseForm()
	if len(r.Form) == 0 {
		if cookie, err := r.Cookie("username"); err != nil {
			mfWebTemplate.ExecuteTemplate(w, "login.html", nil)
			return
		} else {
			ns.Name = cookie.Value
			mfWebTemplate.ExecuteTemplate(w, "chat.html", ns)
			return
		}
	}

	if r.Method == "POST" {
		var user, pass string
		if v, ok := r.Form["username"]; ok && len(v) > 0 {
			user = v[0]
		}
		if v, ok := r.Form["password"]; ok && len(v) > 0 {
			pass = v[0]
		}

		// if !verifyPassword(user, pass) {
		// 	mfWebTemplate.ExecuteTemplate(w, "login.html", nil)
		// 	return
		// }
		ns.Name = user
		if _, ok := r.Form["rememberme"]; ok {
			cookie := http.Cookie{
				Name:  "username",
				Value: user,
			}
			http.SetCookie(w, &cookie)
		}
	}
	mfWebTemplate.ExecuteTemplate(w, "chat.html", ns)
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
