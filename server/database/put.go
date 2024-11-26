package database

import (
	"encoding/json"
	"errors"
	"net/http"
)

func updateDataInDB(query string, data ...any) error {
	db := connect()
	defer db.Close()
	_, err := db.Exec(query, data...)
	if err != nil {
		return err
	}
	return nil
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	query := "UPDATE todos SET task = ?, status = ?, user_id = ?, description = ?, start_time = ?, end_time = ?, image = ?, labels = ? WHERE id = ?"
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
	err = updateDataInDB(
		query,
		todo.Task,
		todo.Status,
		todo.UserID,
		todo.Description,
		todo.StartTime,
		todo.EndTime,
		todo.Image,
		todo.Labels,
		todo.ID,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated Todos!"))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	query := "UPDATE users SET first_name = ?, last_name = ?, username = ?, email = ?, password = ?, image = ? WHERE id = ?"
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
	err = updateDataInDB(
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
		user.Image,
		user.ID,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated Users"))
}

func updatePreferences(w http.ResponseWriter, r *http.Request) {
	query := "UPDATE preferences SET language = ?, theme = ? WHERE user_id = ?"
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
	err = updateDataInDB(
		query,
		preferences.Language,
		preferences.Theme,
		preferences.UserID,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updated Preferences"))
}
