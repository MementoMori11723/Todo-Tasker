package database

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPath           = "todo.db"
	walMode          = "PRAGMA journal_mode=WAL;"
	todosQuery       = "CREATE TABLE todos (id INTEGER PRIMARY KEY, user_id INTEGER, task TEXT, description TEXT, labels TEXT, status TEXT, start_time TEXT, end_time TEXT, image TEXT);"
	usersQuery       = "CREATE TABLE users (id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, username TEXT, email TEXT, password TEXT, image TEXT);"
	preferencesQuery = "CREATE TABLE preferences (user_id INTEGER PRIMARY KEY, theme TEXT, language TEXT);"
	apiRoutes        = map[string]func(http.ResponseWriter, *http.Request){
		"GET /":    getTodos,
		"POST /":   createTodo,
		"PUT /":    updateTodo,
		"DELETE /": deleteTodo,
	}
)

type Todo struct {
	ID          int      `json:"id"`
	UserID      int      `json:"userId"`
	Task        string   `json:"task"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	Status      string   `json:"status"`
	StartTime   string   `json:"startTime"`
	EndTime     string   `json:"endTime"`
	Image       string   `json:"image"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Image     string `json:"image"`
}

type Preferences struct {
	UserID        int    `json:"userId"`
	Theme         string `json:"theme"`
	Language      string `json:"language"`
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
