package database

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPath           = "todo.db"
	walMode          = "PRAGMA journal_mode=WAL;"
	todosQuery       = "CREATE TABLE todos (id INTEGER PRIMARY KEY, user_id INTEGER NOT NULL, task TEXT NOT NULL, description TEXT NOT NULL, labels TEXT, status TEXT NOT NULL, start_time TEXT, end_time TEXT, image TEXT);"
	usersQuery       = "CREATE TABLE users (id INTEGER PRIMARY KEY, first_name TEXT NOT NULL, last_name TEXT NOT NULL, username TEXT NOT NULL, email TEXT NOT NULL, password TEXT NOT NULL, image TEXT);"
	preferencesQuery = "CREATE TABLE preferences (user_id INTEGER PRIMARY KEY, theme TEXT DEFAULT 'light', language TEXT DEFAULT 'en');"

	apiRoutes = map[string]func(http.ResponseWriter, *http.Request){
		"GET /todos":      getTodos,
		"GET /todos/{id}": getTodoByID,
		"POST /todos":     createTodo,
		"PUT /todos":      updateTodo,
		"DELETE /todos":   deleteTodo,

		"GET /users":      getUsers,
		"GET /users/{id}": getUserByID,
		"POST /users":     createUser,
		"PUT /users":      updateUser,
		"DELETE /users":   deleteUser,

		"GET /preferences/{id}": getPreferences,
		"POST /preferences":     createPreferences,
		"PUT /preferences":      updatePreferences,
    "DELETE /preferences":   deletePreferences,
	}
)

type Todo struct {
	ID          int      `json:"id"`
	UserID      int      `json:"userId"`
	Task        string   `json:"task"`
	Description string   `json:"description,omitempty"`
	Labels      []string `json:"labels"`
	Status      string   `json:"status"`
	StartTime   string   `json:"startTime,omitempty"`
	EndTime     string   `json:"endTime,omitempty"`
	Image       string   `json:"image,omitempty"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Image     string `json:"image,omitempty"`
}

type Preferences struct {
	UserID   int    `json:"userId,omitempty"`
	Theme    string `json:"theme,omitempty"`
	Language string `json:"language,omitempty"`
}

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func init() {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Create(dbPath)
		db := connect()
		defer db.Close()
		_, err = db.Exec(
			walMode,
		)
		if err != nil {
			panic(err)
		}
		_, err := db.Exec(
			todosQuery,
		)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(
			usersQuery,
		)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(
			preferencesQuery,
		)
		if err != nil {
			panic(err)
		}
	} else {
		if err != nil {
			panic(err)
		}
	}
}

func errorFunction(w http.ResponseWriter, err error, status int) {
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:      err.Error(),
		StatusCode: status,
	})
}

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	return db
}

func NewApiMux() *http.ServeMux {
	mux := http.NewServeMux()
	for route, handler := range apiRoutes {
		mux.HandleFunc(route, handler)
	}
	return mux
}
