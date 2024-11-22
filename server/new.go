package server

import (
	"embed"
	"net/http"
)

var (
	//go:embed pages/*html
	pages embed.FS

	//go:embed pages/assets/*
	assets embed.FS

	layout = "pages/layout.html"
	routes = map[string]func(http.ResponseWriter, *http.Request){
		"/":      home,
		"/about": about,
		"/todo":  todo,
		"/error": error,
	}
)

func New() *http.ServeMux {
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
	return mux
}
