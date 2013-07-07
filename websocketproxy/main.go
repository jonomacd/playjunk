package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"net/http"
	"net/url"
)

var ConnectedClients map[string]*websocket.Conn = make(map[string]*websocket.Conn)
var addr = flag.String("addr", ":8080", "http service address")

func main() {
	log.Println("hello world")

	flag.Parse()

	//http.HandleFunc("/", homeHandler)
	http.Handle("/ws", websocket.Handler(initialize))
	http.HandleFunc("/forward", forward)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
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
	}

	go func(sock *websocket.Conn) {
		for {
			var message string
			err = websocket.Message.Receive(sock, &message)

			if err != nil {
				log.Println(err)
				_, err := http.PostForm("http://localhost:8099/delete",
					url.Values{"id": {u4.String()}})
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
	}(ws)
}

func forward(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	body := r.FormValue("body")

	ConnectedClients[id].Write([]byte(body))
}
