package handler

import (
	"encoding/json"
	"go-server/database"
	"log"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	Routes = map[string]func(http.ResponseWriter, *http.Request) {
    "GET /tasks": GetTask,
    "POST /tasks": InsertTask,
    "PUT /tasks": UpdateTask,
    "DELETE /tasks": DeleteTask,
  }
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Get(
		"SELECT * FROM tasks WHERE userId =" +
			r.URL.Query().Get("userId"),
	)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	var tasks []Task
	for rows.Next() {
		var task Task
		rows.Scan(
      &task.ID, 
      &task.UserID, 
      &task.TaskName, 
      &task.Description,
    )
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
}

func InsertTask(w http.ResponseWriter, r *http.Request) {
  var task Task
  json.NewDecoder(r.Body).Decode(&task)
  err := database.Insert(
    "INSERT INTO tasks (userId, taskName, description) VALUES (?, ?, ?)",
    []string{
      string(task.UserID),
      task.TaskName,
      task.Description,
    },
  )
  if err != nil {
    log.Println(err)
  }
  json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
  var task Task
  json.NewDecoder(r.Body).Decode(&task)
  err := database.Update(
    "UPDATE tasks SET taskName = ?, description = ? WHERE id = ?",
    []string{
      task.TaskName,
      task.Description,
      string(task.ID),
    },
  )
  if err != nil {
    log.Println(err)
  }
  json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
  var task Task
  json.NewDecoder(r.Body).Decode(&task)
  err := database.Delete(
    "DELETE FROM tasks WHERE id = ?",
    []string{
      string(task.ID),
    },
  )
  if err != nil {
    log.Println(err)
  }
  json.NewEncoder(w).Encode(task)
}
