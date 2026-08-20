package handlers

import (
	"github.com/go-chi/chi/v5"
)

var Main_router = chi.NewRouter()

func RegisterHandlers() {
	Main_router.Get("/{short_code}", getShorten)
	Main_router.Post("/", createShorten)
}
