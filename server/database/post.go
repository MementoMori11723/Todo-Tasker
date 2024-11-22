package database

import (
  "net/http"
)

func createTodo(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("This is a POST request"))
} 
