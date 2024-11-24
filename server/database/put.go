package database

import (
  "net/http"
)

func updateTodo(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("This is a PUT request"))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("This is a PUT request"))
}

func updatePreferences(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("This is a PUT request"))
}
