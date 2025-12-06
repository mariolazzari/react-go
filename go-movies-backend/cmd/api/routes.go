package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	// create new router
	mux := chi.NewRouter()
	// middlewares
	mux.Use(middleware.Recoverer)
	// routes
	mux.Get("/", app.Home)

	return mux
}
