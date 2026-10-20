package service

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/vatut007/short_link_increment7/cmd/shortener/config"
	"github.com/vatut007/short_link_increment7/store"
	"github.com/vatut007/short_link_increment7/utils"
)

type UrlService struct {
	Url            string
	Request        *http.Request
	OriginalURL    string
	ResponseWriter http.ResponseWriter
	CreateUrlInter URLCreateInterface
	shortURL       string
}

func (u *UrlService) getFabByContentType(content_type ContentType) error {
	if !content_type.isValid() {
		return ErrNoType
	}
	switch content_type {
	case ContentText:
		u.CreateUrlInter = &UrlCreateText{UrlService: u}
	case ContentJson:
		u.CreateUrlInter = &UrlCreateJson{UrlService: u}
	}
	return nil
}

func (ui *UrlService) GenerateShortUrl(contentType string) error {
	// r := ui.Request
	// content_type := r.Header.Get("Content-Type")
	if err := ui.getFabByContentType(ContentType(contentType)); err != nil {
		return err
	}
	err := ui.CreateUrlInter.loadBody()
	if err != nil {
		return err
	}
	originalURL := ui.OriginalURL
	if strings.TrimSpace(originalURL) == "" {
		return ErrEmptyUrl
	}
	if _, err := url.ParseRequestURI(originalURL); err != nil {
		return ErrInvalidUrl
	}
	code := utils.GenerateShortCodeUrl(originalURL)
	ui.shortURL = config.ConfigApp.BaseURL + "/" + code
	store.Store.Store(code, originalURL)
	return nil
}

func (ui *UrlService) GetHeaderParams() (string, string) {
	key, value := ui.CreateUrlInter.getHeader()
	return key, value
}

func (ui *UrlService) WriteBody(w http.ResponseWriter) {
	uc := ui.CreateUrlInter
	uc.writeBody(w)
}
