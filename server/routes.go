package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *Config) Routes() http.Handler {
	// new mux instance
	mux := chi.NewRouter()
	// middleware configuration
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://", "https://"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "x-CSRF-toke"},
	}))

	mux.Post("/password", app.addNewPassword)
	mux.Get("/password", app.GetPasswordList)
	mux.Get("/password/{key}", app.GetPlainPassword)
	mux.Get("/", app.RenderUI)
	return mux
}
