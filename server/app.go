package main

import (
	"fmt"
	"go-server/config"
	"go-server/handler"
	"net/http"
)

func main() {
	for route, function := range handler.Routes {
		http.HandleFunc(route, function)
	}
	fmt.Println(
		"Server" +
			" running " +
			" on " +
			" port " +
			" http://localhost" +
			config.Port,
	)
	http.ListenAndServe(config.Port, nil)
}
