package service

import (
	"net/http"
)

func GetHttpErr(err error, w http.ResponseWriter) {
	switch err {
	case ErrEmptyUrl:
		err_message := ErrEmptyUrl.Error()
		http.Error(w, err_message, http.StatusBadRequest)
		return
	case ErrNoType:
		err_message := ErrNoType.Error()
		http.Error(w, err_message, http.StatusBadRequest)
		return
	case ErrInvalidUrl:
		err_message := ErrInvalidUrl.Error()
		http.Error(w, err_message, http.StatusBadRequest)
	}
}
