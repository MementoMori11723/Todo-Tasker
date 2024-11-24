package server

import (
	"Todo-Tasker/server/database"
	"Todo-Tasker/server/middleware"
	"embed"
	"net/http"
)

var (
	//go:embed pages/*html
	pages embed.FS

	//go:embed pages/assets/*
	assets embed.FS

	layout = "pages/layout.html"
  api = database.NewApiMux()
	routes = map[string]func(http.ResponseWriter, *http.Request){
		"/":      home,
		"/about": about,
		"/todo":  todo,
    "/login": login,
		"/error": error,
	}
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/assets/",
		http.StripPrefix(
			"/assets",
			http.FileServer(
				http.Dir(
					"server/pages/assets",
				),
			),
		),
	)
	for path, handler := range routes {
		mux.HandleFunc(path, handler)
	}
  mux.Handle("/api/", 
    http.StripPrefix("/api", api),
  )
  newMux := middleware.Log(mux)
	return newMux
}
