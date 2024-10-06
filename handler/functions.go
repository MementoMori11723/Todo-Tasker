package handler

import (
	"net/http"
)



var (
	Routes = make(map[string]func(http.ResponseWriter, *http.Request))
  // {
  //   "GET /tasks": GetTask,
  //   "POST /tasks": InsertTask,
  //   "PUT /tasks": UpdateTask,
  //   "DELETE /tasks": DeleteTask,
  // }
)
