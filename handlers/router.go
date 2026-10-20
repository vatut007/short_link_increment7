package handlers

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/", createShorten)
	r.Get("/{short_code}", getShorten)
	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", apiCreateShorten)
	})
	return r
}
