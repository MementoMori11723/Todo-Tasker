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
		log.Println("Server started on http://localhost:8000")
		log.Println("Press Enter to stop the server")
		if err := http.ListenAndServe(":8000", server); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Scanln()
}
