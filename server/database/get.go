package database

import (
	"encoding/json"
	"errors"
	"net/http"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "" || r.Header.Get("Authorization") != "Bearer token" {
		errorFunction(w, errors.New("Unauthorized"), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT id, task, status, user_id, description, start_time, end_time, image, labels FROM todos"
	var todos []Todo
  db := connect()	
  defer db.Close()
  rows, err := db.Query(query)
  for rows.Next() {
    var todo Todo
    err := rows.Scan(
      &todo.ID, 
      &todo.Task, 
      &todo.Status, 
      &todo.UserID, 
      &todo.Description, 
      &todo.StartTime, 
      &todo.EndTime, 
      &todo.Image, 
      &todo.Labels,
    )
    if err != nil {
      errorFunction(w, err, http.StatusInternalServerError)
      return
    }
    todos = append(todos, todo)
  }
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	if todos[0].ID == 0 {
		errorFunction(w, errors.New("No Todos Found!"), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "" || r.Header.Get("Authorization") != "Bearer token" {
		errorFunction(w, errors.New("Unauthorized"), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT id, first_name, last_name, username, email, password, image FROM users"
	var users []User
  db := connect()
  defer db.Close()
  rows, err := db.Query(query)
  for rows.Next() {
    var user User
    err := rows.Scan(
      &user.ID, 
      &user.FirstName, 
      &user.LastName, 
      &user.Username, 
      &user.Email, 
      &user.Password, 
      &user.Image,
    )
    if err != nil {
      errorFunction(w, err, http.StatusInternalServerError)
      return
    }
    users = append(users, user)
  }
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	if users[0].ID == 0 {
		errorFunction(w, errors.New("No User Found!"), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT id, first_name, last_name, username, email, password, image FROM users WHERE id = " + r.PathValue("id") + ";"
	var user User
  db := connect()
  defer db.Close()
  err := db.QueryRow(query).Scan(
    &user.ID, 
    &user.FirstName, 
    &user.LastName, 
    &user.Username, 
    &user.Email, 
    &user.Password, 
    &user.Image,
  )
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	if user.ID == 0 {
		errorFunction(w, errors.New("No User Found!"), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
}

func getPreferences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT user_id, theme, language FROM preferences WHERE user_id = " + r.PathValue("id") + ";"
	var preference Preferences
  db := connect()
	err := db.QueryRow(query).Scan(
		&preference.UserID,
		&preference.Theme,
		&preference.Language,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	if preference.UserID == 0 {
		errorFunction(w, errors.New("No Preferences Found!"), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(preference); err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT id, task, status, user_id, description, start_time, end_time, image, labels FROM todos WHERE id = " + r.PathValue("id") + ";"
	var todo Todo
  db := connect()
	err := db.QueryRow(query).Scan(
		&todo.ID,
		&todo.Task,
		&todo.Status,
		&todo.UserID,
		&todo.Description,
		&todo.StartTime,
		&todo.EndTime,
		&todo.Image,
		&todo.Labels,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	if todo.ID == 0 {
		errorFunction(w, errors.New("No Todos Found!"), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
}
