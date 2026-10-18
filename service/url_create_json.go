package service

import (
	"encoding/json"
	"net/http"

	"github.com/vatut007/short_link_increment7/dto"
)

type UrlCreateJson struct {
	UrlService *UrlService
	urlBody    dto.UrlCreateDto
}

func (ui *UrlCreateJson) loadBody() error {
	var urlBody dto.UrlCreateDto
	body := ui.UrlService.Request.Body
	if err := json.NewDecoder(body).Decode(&urlBody); err != nil {
		return err
	}
	ui.urlBody = urlBody
	ui.saveOriginalUrl(urlBody.URL)
	return nil
}

func (ui *UrlCreateJson) getHeader() (string, string) {
	return "Content-Type", "application/json"
}

func (ui UrlCreateJson) writeBody(w http.ResponseWriter) {

}

func (ui *UrlCreateJson) saveOriginalUrl(url string) {
	ui.UrlService.OriginalURL = url
}
