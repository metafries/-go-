package mfportal

import (
	"-go-/MF/dblayer"
	"-go-/MF/mfweb/mfrestapi"
	"-go-/_mfConfig"
	"bufio"
	"html/template"
	"log"
	"net"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

var mfWebTemplate *template.Template
var historyLog = struct {
	logs []string
	*sync.RWMutex
}{RWMutex: new(sync.RWMutex)}

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
	http.Handle("/chatRoom/", websocket.Handler(chatWS))
	go func() {
		err = http.ListenAndServeTLS(":8062", "cert.pem", "key.pem", nil)
		log.Println(err)
	}()
	err = http.ListenAndServe(":8061", nil)
	return err
}

func chatWS(ws *websocket.Conn) {
	conn, err := net.Dial("tcp", "127.0.0.1:2100")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	historyLog.RLock()
	for _, log := range historyLog.logs {
		err := websocket.Message.Send(ws, log)
		if err != nil {
			historyLog.RUnlock()
			return
		}
	}
	historyLog.RUnlock()

	go func() {
		// Use scanner to receive chat messages
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			err := websocket.Message.Send(ws, message)
			if err != nil {
				return
			}
		}
	}()

	for {
		// Receive text frame
		var msg string
		err := websocket.Message.Send(ws, &msg)
		if err != nil {
			return
		}
		_, err = conn.Write([]byte(msg))
		if err == nil {
			historyLog.Lock()
			if len(historyLog.logs) > 20 {
				historyLog.logs = historyLog.logs[1:]
			}
			historyLog.logs = append(historyLog.logs, msg)
			historyLog.Unlock()
		}
	}
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

		if !verifyPassword(user, pass) {
			mfWebTemplate.ExecuteTemplate(w, "login.html", nil)
			return
		}
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

func verifyPassword(username, password string) bool {
	return true
}
