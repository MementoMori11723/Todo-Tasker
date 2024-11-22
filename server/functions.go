package server

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(
      w, r, "/error", 
      http.StatusFound,
    )
	}
  tmpl, err := template.ParseFS(
    pages, layout, 
    "pages/home.html",
  )	
  if err != nil {
    http.Error(
      w, err.Error(), 
      http.StatusInternalServerError,
    )
    return
  }
  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(
      w, err.Error(), 
      http.StatusInternalServerError,
    )
    return
  }
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Redirect(w, r, "/error", http.StatusFound)
	}
  tmpl, err := template.ParseFS(
    pages, layout, 
    "pages/about.html",
  )	
  if err != nil {
    http.Error(
      w, err.Error(), 
      http.StatusInternalServerError,
    )
    return
  }
  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(
      w, err.Error(), 
      http.StatusInternalServerError,
    )
    return
  }
}

func todo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/todo" {
		http.Redirect(w, r, "/error", http.StatusFound)
	}
  tmpl, err := template.ParseFS(
    pages, layout, 
    "pages/todo.html",
  )	
  if err != nil {
    http.Error(
      w, err.Error(), 
      http.StatusInternalServerError,
    )
    return
  }
  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(
      w, err.Error(), 
      http.StatusInternalServerError,
    )
    return
  }
}

func error(w http.ResponseWriter, r *http.Request) {
  tmpl, err := template.ParseFS(
    pages, layout, 
    "pages/error.html",
  )	
  if err != nil {
    http.Error(
      w, err.Error(), 
      http.StatusInternalServerError,
    )
    return
  }
  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(
      w, err.Error(), 
      http.StatusInternalServerError,
    )
    return
  }
}
