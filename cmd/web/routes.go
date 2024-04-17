package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	r := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	r.HandleFunc("/", app.home)
	r.HandleFunc("/snippet/view", app.snippetView)
	r.HandleFunc("/snippet/create", app.snippetCreate)

	return app.logRequest(secureHeaders(r))
}
