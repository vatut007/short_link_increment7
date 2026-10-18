package service

import (
	"io"
	"net/http"
)

type UrlCreateText struct {
	UrlService *UrlService
	BodyBytes  []byte
}

func (ui *UrlCreateText) loadBody() error {
	bodyBytes, err := io.ReadAll(ui.UrlService.Request.Body)
	if err != nil {
		return err
	}
	ui.BodyBytes = bodyBytes
	ui.saveOriginalUrl(string(bodyBytes))
	return nil
}

func (ui *UrlCreateText) getHeader() (string, string) {
	return "Content-Type", "text/plain"
}

func (ui UrlCreateText) writeBody(w http.ResponseWriter) {
	if _, err := w.Write([]byte(ui.UrlService.shortURL)); err != nil {
		panic(err)
	}
}
func (ui *UrlCreateText) saveOriginalUrl(url string) {
	ui.UrlService.OriginalURL = url
}
