package database

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPath     = "todo.db"
	todosQuery = "CREATE TABLE todos (id INTEGER PRIMARY KEY, user_id INTEGER, task TEXT, description TEXT, labels TEXT);"
	usersQuery = "CREATE TABLE users (id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, username TEXT, email TEXT, password TEXT, image TEXT);"
	apiRoutes  = map[string]func(http.ResponseWriter, *http.Request){
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

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func init() {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Create(dbPath)
		db := connect()
		defer db.Close()
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
