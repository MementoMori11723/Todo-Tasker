package main

import (
	"go-server/config"
	"go-server/handler"
	"net/http"
	"log"
)

func main() {
	for route, function := range handler.Routes {
		http.HandleFunc(route, function)
	}
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
  })
	log.Println("Server started on port", config.Port)
	http.ListenAndServe(config.Port, nil)
}
