package database

import (
	"encoding/json"
	"errors"
	"net/http"
)

func createDataInDB(query string, data ...any) error {
	db := connect()
	defer db.Close()
	_, err := db.Exec(query, data...)
	if err != nil {
		return err
	}
	return nil
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	query := "INSERT INTO todos (task, status, user_id, description, start_time, end_time, image, labels) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	var todo Todo
	if r.Body == nil {
		errorFunction(w, errors.New("Please send a request body"), http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		errorFunction(w, err, http.StatusBadRequest)
		return
	}
	err = createDataInDB(
		query,
		todo.Task,
		todo.Status,
		todo.UserID,
		todo.Description,
		todo.StartTime,
		todo.EndTime,
		todo.Image,
		todo.Labels,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
  w.Write([]byte("Todo created successfully"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
  query := "INSERT INTO users (first_name, last_name, username, email, password, image) VALUES (?, ?, ?, ?, ?, ?)"
  var user User
  if r.Body == nil {
    errorFunction(w, errors.New("Please send a request body"), http.StatusBadRequest)
    return
  }
  err := json.NewDecoder(r.Body).Decode(&user)
  if err != nil {
    errorFunction(w, err, http.StatusBadRequest)
    return
  }
  err = createDataInDB(
    query,
    user.FirstName,
    user.LastName,
    user.Username,
    user.Email,
    user.Password,
    user.Image,
  )
  if err != nil {
    errorFunction(w, err, http.StatusInternalServerError)
    return
  }
  w.Write([]byte("User created successfully"))
}

func createPreferences(w http.ResponseWriter, r *http.Request) {
  query := "INSERT INTO preferences (user_id, theme, language) VALUES (?, ?, ?)"
  var preferences Preferences
  if r.Body == nil {
    errorFunction(w, errors.New("Please send a request body"), http.StatusBadRequest)
    return
  }
  err := json.NewDecoder(r.Body).Decode(&preferences)
  if err != nil {
    errorFunction(w, err, http.StatusBadRequest)
    return
  }
  err = createDataInDB(
    query,
    preferences.UserID,
    preferences.Theme,
    preferences.Language,
  )
  if err != nil {
    errorFunction(w, err, http.StatusInternalServerError)
    return
  }
  w.Write([]byte("Preferences created successfully"))
}
