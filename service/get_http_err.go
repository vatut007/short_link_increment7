package service

import (
	"errors"
	"net/http"
)

func GetHttpErr(err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, ErrEmptyUrl), errors.Is(err, ErrInvalidUrl), errors.Is(err, ErrNoType):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}
