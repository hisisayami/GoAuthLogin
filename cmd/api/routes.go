package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	//create a router mux
	mux := chi.NewRouter()

	//apply middleware
	mux.Use(app.enableCORS)

	//mux.Get("/", app.CreateUserTable)
	mux.Post("/authenticate", app.authenticate)
	mux.Post("/createUser", app.createUser)


	return mux

}
