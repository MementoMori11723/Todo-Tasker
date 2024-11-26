package database

import (
	"encoding/json"
	"errors"
	"net/http"
)

func getDataFromDB(query string, data ...any) error {
	db := connect()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(data...)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "" || r.Header.Get("Authorization") != "Bearer token" {
		errorFunction(w, errors.New("Unauthorized"), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT id, task, status, user_id, description, start_time, end_time, image, labels FROM todos"
	var todo Todo
	err := getDataFromDB(
		query,
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "" || r.Header.Get("Authorization") != "Bearer token" {
		errorFunction(w, errors.New("Unauthorized"), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT id, first_name, last_name, username, email, password, image FROM users"
	var user User
	err := getDataFromDB(
		query,
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Image,
		&user.LastName,
		&user.Username,
		&user.FirstName,
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

func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := "SELECT id, first_name, last_name, username, email, password, image FROM users WHERE id = " + r.PathValue("id") + ";"
	var user User
	err := getDataFromDB(
		query,
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Image,
		&user.LastName,
		&user.Username,
		&user.FirstName,
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
	err := getDataFromDB(
		query,
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
	err := getDataFromDB(
		query,
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
