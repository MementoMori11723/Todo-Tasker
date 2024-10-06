package database

import ()

type Task struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
}


