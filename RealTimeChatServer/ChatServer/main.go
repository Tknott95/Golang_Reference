package main

import (
	"html/template"
	"log"
	"net/http"

	hub "github.com/tknott95/Golang_Reference/RealTimeChatServer/ChatServer/Controllers/hubCtrl"
)

var index = template.Must(template.ParseFiles("./Views/index.html"))

func home(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
}

func main() {
	go hub.DefaultHub.Start()

	log.Println(`Chat server launched on port 8080 🚀`)

	http.HandleFunc("/", home)
	http.HandleFunc("/ws", hub.WSHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
