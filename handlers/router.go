package handlers

import (
	"github.com/go-chi/chi/v5"
)

var Main_router = chi.NewRouter()
var apiRouter = chi.NewRouter()

func RegisterHandlers() {
	Main_router.Get("/{short_code}", getShorten)
	Main_router.Route("/api", func(r chi.Router) {
		r.Post("/shorten", apiCreateShorten)
	})
	Main_router.Post("/", createShorten)
}

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/", createShorten)
	r.Get("/{short_code}", getShorten)
	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", apiCreateShorten)
	})
	return r
}
