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
	log.Println("Server started on port", config.Port)
	http.ListenAndServe(config.Port, nil)
}
