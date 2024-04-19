package main

import (
	"net/http"
	"github.com/go-chi/chi"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Get("/", app.home)
	r.Get("/snippet/view/{id}", app.snippetView)
	r.Get("/snippet/create", app.snippetCreate)
	r.Post("/snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(r)
}
