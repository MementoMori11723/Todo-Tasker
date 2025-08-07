package server

import (
	"embed"
	"html/template"
	"log/slog"
	"net/http"
)

type (
	routes map[string]http.HandlerFunc

	errorType struct {
		Title string
		Msg   string
	}
)

var (
	//go:embed pages
	_pages embed.FS

	// //go:embed assets
	// _assets embed.FS

	_dir    = "pages/"
	_layout = _dir + "layout.html"

	_indexTempl *template.Template
	_aboutTempl *template.Template
	_loginTempl *template.Template
	_errorTempl *template.Template

	_routes = routes{
		"/":       indexPage,
		"/404/":   pageNotFound,
		"/about/": aboutPage,
		"/error/": errorPage,
	}
)

func init() {
	_indexTempl = getTemplate("index.html")
	_aboutTempl = getTemplate("about.html")
	_loginTempl = getTemplate("login.html")
	_errorTempl = getTemplate("error.html")
}

func getTemplate(filename string) *template.Template {
	temp, err := template.ParseFS(_pages, _layout, _dir+filename)
	if err != nil {
		slog.Error(err.Error())
	}
	return temp
}

func New() *http.ServeMux {
	_mux := http.NewServeMux()
	// _mux.Handle("/assets/", http.FileServerFS(_assets))
	for _route, _handler := range _routes {
		_mux.HandleFunc(
			_route, _handler,
		)
	}
	return _mux
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusFound)
	}
	if err := _indexTempl.Execute(w, nil); err != nil {
		slog.Error(err.Error())
		http.Redirect(w, r, "/404", http.StatusFound)
	}
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	if err := _aboutTempl.Execute(w, nil); err != nil {
		slog.Error(err.Error())
		http.Redirect(w, r, "/404", http.StatusFound)
	}
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := _errorTempl.Execute(w, errorType{
		Title: "404",
		Msg:   "Page Not found!",
	}); err != nil {
		slog.Error(err.Error())
		http.Redirect(w, r, "/error", http.StatusFound)
	}
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	if err := _errorTempl.Execute(w, errorType{
		Title: "500",
		Msg:   "Oops! somthing went Wrong!",
	}); err != nil {
		http.Error(w, "Internal error!", http.StatusInternalServerError)
		slog.Error(err.Error())
	}
}

