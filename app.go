package main

import (
	"Todo-Tasker/server"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
  PORT := flag.String("port", "8000", "Port to run the server on")
  flag.Parse()
	go func() {
		server := server.New()
    mux := http.Server{
      Addr:    ":" + *PORT,
      Handler: server,
    }
		log.Println("Server started on http://localhost:" + *PORT)
		log.Println("Press Enter to stop the server")
    if err := mux.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Scanln()
}
