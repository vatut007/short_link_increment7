package service

import "net/http"

type URLCreateInterface interface {
	loadBody() error
	getHeader() (string, string)
	writeBody(http.ResponseWriter)
	saveOriginalUrl(url string)
}
