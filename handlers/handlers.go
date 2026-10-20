package handlers

import (
	"mime"
	"net/http"

	"github.com/vatut007/short_link_increment7/dto"
	"github.com/vatut007/short_link_increment7/service"
	"github.com/vatut007/short_link_increment7/store"

	"github.com/go-chi/chi/v5"
)

type UrlUserCreate = dto.UrlCreateDto

func createShorten(w http.ResponseWriter, r *http.Request) {
	content_type := r.Header.Get("Content-Type")
	mt, _, err := mime.ParseMediaType(content_type)
	if err != nil || mt != "text/plain" {
		http.Error(w, "Content type must be text/plain", http.StatusUnsupportedMediaType)
		return
	}
	urlSrv := service.UrlService{
		Request:        r,
		ResponseWriter: w,
	}
	err = urlSrv.GenerateShortUrl(mt)
	if err != nil {
		service.GetHttpErr(err, w)
		return
	}
	w.Header().Set(urlSrv.GetHeaderParams())
	w.WriteHeader(http.StatusCreated)
	urlSrv.WriteBody(w)
}

func apiCreateShorten(w http.ResponseWriter, r *http.Request) {
	content_type := r.Header.Get("Content-Type")
	mt, _, err := mime.ParseMediaType(content_type)
	if err != nil || mt != "application/json" {
		http.Error(w, "Content type must be application/json", http.StatusUnsupportedMediaType)
		return
	}
	urlSrv := service.UrlService{
		Request:        r,
		ResponseWriter: w,
	}
	if err := urlSrv.GenerateShortUrl(mt); err != nil {
		service.GetHttpErr(err, w)
		return
	}
	w.Header().Set(urlSrv.GetHeaderParams())
	w.WriteHeader(http.StatusCreated)
	urlSrv.WriteBody(w)

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
