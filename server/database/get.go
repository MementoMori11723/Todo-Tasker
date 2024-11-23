package database

import (
	"encoding/json"
	"net/http"
)

var query = "SELECT * FROM users"

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := connect()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:      err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	defer rows.Close()
	var user User
	for rows.Next() {
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
      w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Error:      err.Error(),
				StatusCode: http.StatusInternalServerError,
			})
			return
		}
	}

	if user.ID == 0 {
    w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:      "No user found",
			StatusCode: http.StatusNotFound,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
    w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error:      err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
}
