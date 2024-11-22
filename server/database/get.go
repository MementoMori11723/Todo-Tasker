package database

import "net/http"

func getTodos(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("This is a GET request"))
}
