package handlers

import (
	"io"
	"net/http"
	"net/url"
	"shortener/cmd/shortener/config"
	"shortener/store"
	"shortener/utils"
	"strings"

	"github.com/go-chi/chi/v5"
)

func createShorten(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "text/plain" {
		http.Error(w, "Content type must be text/plain", http.StatusUnsupportedMediaType)
		return
	}
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	originalURL := string(bodyBytes)
	if strings.TrimSpace(originalURL) == "" {
		http.Error(w, "Url is empty", http.StatusBadRequest)
		return
	}
	if _, err := url.ParseRequestURI(originalURL); err != nil {
		http.Error(w, "Invalid url", http.StatusBadRequest)
		return
	}
	code := utils.GenerateShortCodeUrl(originalURL)
	shortURL := config.ConfigApp.BaseURL + "/" + code
	store.Store.Store(code, originalURL)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}

func getShorten(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "short_code")
	val, exists := store.Store.Load(code)
	if !exists {
		http.Error(w, "Short url not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, val.(string), http.StatusTemporaryRedirect)
}
