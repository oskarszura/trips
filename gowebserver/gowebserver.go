package gowebserver

import (
	"fmt"
	"net/http"
	"log"
)

type WebServer struct {
	Router		UrlRouter
}

func (s *WebServer) RunServer(port string) {
	staticFileServer := http.FileServer(http.Dir("public"))

	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))
	http.HandleFunc("/", s.Router.route)

	fmt.Println("Setting up server on " + port + " port")
	fmt.Println("Listening...")

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}