package database

import (
  "net/http"
)

func deleteTodo(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("This is a DELETE request"))
}
