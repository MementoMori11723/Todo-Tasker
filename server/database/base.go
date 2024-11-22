package database

import "net/http"

var apiRoutes = map[string]func(http.ResponseWriter, *http.Request){
	"GET /":    getTodos,
	"POST /":   createTodo,
	"PUT /":    updateTodo,
	"DELETE /": deleteTodo,
}

func NewApiMux() *http.ServeMux {
  mux := http.NewServeMux()
  for route, handler := range apiRoutes {
    mux.HandleFunc(route, handler)
  }
  return mux
}
