package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var ConnectedClients map[string]*websocket.Conn = make(map[string]*websocket.Conn)
var addr = flag.String("addr", ":8080", "http service address")
var homeTempl *template.Template

func main() {
	log.Println("hello world")
	defer removeAllUsers()
	rootDir := flag.String("dir", "../resources/", "resource directory")

	flag.Parse()

	homeTempl = template.Must(template.ParseFiles(*rootDir + "canvasPage.html"))
	http.Handle("/inc/", http.StripPrefix("/inc/", http.FileServer(http.Dir(*rootDir))))
	http.HandleFunc("/", homeHandler)
	http.Handle("/ws", websocket.Handler(initialize))
	http.HandleFunc("/forward", forward)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func initialize(ws *websocket.Conn) {
	u4, err := uuid.NewV4()
	ConnectedClients[u4.String()] = ws
	if err != nil {
		log.Println("error id'ing ::", err)
	}
	_, err = http.PostForm("http://localhost:8099/new",
		url.Values{"id": {u4.String()}})
	if err != nil {
		log.Println("error posting ::", err)
		return
	}

	for {
		var message string
		err = websocket.Message.Receive(ws, &message)

		if err != nil {
			log.Println("receive on websocket error:", err)
			err := removeUser(u4.String())
			if err != nil {
				log.Println("error posting ::", err)
			}
			break
		}

		_, err = http.PostForm("http://localhost:8099/data",
			url.Values{"id": {u4.String()}, "body": {message}})
		if err != nil {
			log.Println("error posting ::", err)
		}
	}
}

func forward(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	body := r.FormValue("body")
	log.Println("the body", body)
	ConnectedClients[id].Write([]byte(body))
}

func removeUser(id string) error {
	_, err := http.PostForm("http://localhost:8099/delete",
		url.Values{"id": {id}})
	return err
}

func removeAllUsers() {
	for id, _ := range ConnectedClients {
		removeUser(id)
	}
}
