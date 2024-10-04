package main

import (
	"fmt"
	"go-server/database"
	"net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    conn := database.Connect()
    conn.Close()
    fmt.Fprintf(w, "Hello, World!")
  })

  http.ListenAndServe(":8080", nil)
}
