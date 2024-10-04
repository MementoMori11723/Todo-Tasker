package handler

import (
	"fmt"
	"go-server/database"
	"net/http"
)

var (
  Routes = map[string]func(http.ResponseWriter, *http.Request){
    "GET /tasks": getAll,
    "GET /tasks/{id}": getOne,
    "POST /tasks": postTask,
    "PUT /tasks": putTask,
    "DELETE /tasks": deleteTask,
  }
)

// need to work on these functions... 

func getAll(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	fmt.Fprintf(w, "All tasks")
}

func getOne(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks WHERE id = ?", 1)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	fmt.Fprintf(w, "One task")
}

func postTask(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	_, err := db.Exec("INSERT INTO tasks (name, description) VALUES (?, ?)", "Task 1", "Description 1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "Task created")
}

func putTask(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	_, err := db.Exec("UPDATE tasks SET name = ?, description = ? WHERE id = ?", "Task 1 updated", "Description 1 updated", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "Task updated")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	db := database.Connect()
	defer db.Close()

	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "Task deleted")
}
