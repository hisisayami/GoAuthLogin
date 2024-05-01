package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	//create a router mux
	mux := chi.NewRouter()

	mux.Get("/", app.CreateUserTable)
	mux.Get("/hello", app.Hello)

	return mux

}
