package database

import (
	"encoding/json"
	"errors"
	"net/http"
)

func deleteDataInDB(query string, data ...any) error {
	db := connect()
	defer db.Close()
	_, err := db.Exec(query, data...)
	if err != nil {
		return err
	}
	return nil
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	query := "DELETE FROM todos WHERE id = ?"
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
	err = deleteDataInDB(
		query,
		todo.ID,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	w.Write(
		[]byte("Deleted Todo of ID: " +
			string(
				rune(todo.ID),
			),
		),
	)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	query := "DELETE FROM users WHERE id = ?"
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
	err = deleteDataInDB(
		query,
		user.ID,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	w.Write(
		[]byte("Deleted User of ID: " +
			string(
				rune(user.ID),
			),
		),
	)
}

func deletePreferences(w http.ResponseWriter, r *http.Request) {
	query := "DELETE FROM preferences WHERE user_id = ?"
	var preference Preferences
	if r.Body == nil {
		errorFunction(w, errors.New("Please send a request body"), http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&preference)
	if err != nil {
		errorFunction(w, err, http.StatusBadRequest)
		return
	}
	err = deleteDataInDB(
		query,
		preference.UserID,
	)
	if err != nil {
		errorFunction(w, err, http.StatusInternalServerError)
		return
	}
	w.Write(
		[]byte("Deleted Preferences of User ID: " +
			string(
				rune(preference.UserID),
			),
		),
	)
}
