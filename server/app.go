package main

import (
	"fmt"
	"go-server/handler"
	"net/http"
)

func main() {
	for route, function := range handler.Routes {
		http.HandleFunc(route, function)
	}
	fmt.Println("Server running on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
