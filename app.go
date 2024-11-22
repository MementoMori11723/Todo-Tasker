package main

import (
	"Todo-Tasker/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	go func() {
		server := server.New()
    mux := http.Server{
      Addr:    ":8000",
      Handler: server,
    }
		log.Println("Server started on http://localhost:8000")
		log.Println("Press Enter to stop the server")
    if err := mux.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Scanln()
}
