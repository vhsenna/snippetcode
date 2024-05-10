package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	}))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Get("/ping", ping)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	r.Method(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	r.Method(http.MethodGet, "/snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	r.Method(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
	r.Method(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	r.Method(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
	r.Method(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	r.Method(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	r.Method(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
	r.Method(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(r)
}
