package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vatut007/short_link_increment7/cmd/shortener/config"
	"github.com/vatut007/short_link_increment7/dto"
)

func Test_create_handlers(t *testing.T) {
	type want struct {
		code        int
		contentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "test create shorten",
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config.ConfigApp.BaseURL = "http://example.com"
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://www.yandex.ru"))
			request.Header.Set("Content-Type", test.want.contentType)
			w := httptest.NewRecorder()
			createShorten(w, request)
			res := w.Result()
			assert.Equal(t, test.want.code, res.StatusCode)
			defer func() {
				if err := res.Body.Close(); err != nil {
					log.Printf("close body: %v", err)
				}
			}()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Regexp(t, `^http://example\.com/[A-Za-z0-9]{8}$`, string(resBody))
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
func Test_create_api_handlers(t *testing.T) {
	type want struct {
		code        int
		contentType string
		body        dto.UrlCreateDto
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "test create shorten api",
			want: want{
				code:        201,
				contentType: "application/json",
				body:        dto.UrlCreateDto{URL: "https://www.yandex.ru"},
			},
		},
		{
			name: "test create shorten wrong content type",
			want: want{
				code:        415,
				contentType: "text/plain",
				body:        dto.UrlCreateDto{URL: "https://www.yandex.ru"},
			},
		},
		{
			name: "test empty url",
			want: want{
				code:        400,
				contentType: "application/json",
				body:        dto.UrlCreateDto{},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config.ConfigApp.BaseURL = "http://example.com"
			// body_dto := dto.UrlCreateDto{URL: "https://www.yandex.ru"}
			reqBody, err := json.Marshal(test.want.body)
			require.NoError(t, err)
			request := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(string(reqBody)))
			request.Header.Set("Content-Type", test.want.contentType)
			w := httptest.NewRecorder()
			router := NewRouter()
			router.ServeHTTP(w, request)
			res := w.Result()
			defer func() {
				if err := res.Body.Close(); err != nil {
					log.Printf("close body err:%v", err)
				}
			}()
			assert.Equal(t, test.want.code, res.StatusCode)
			if test.want.code == http.StatusCreated {
				assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
			}
		})
	}
}
func Test_get_handlers(t *testing.T) {
	type want struct {
		code        int
		contentType string
		location    string
	}
	tests := []struct {
		name        string
		originalURL string
		want        want
	}{
		{
			name:        "test get original link",
			originalURL: "https://www.yandex.ru",
			want: want{
				code:        http.StatusTemporaryRedirect,
				contentType: "text/plain",
				location:    "https://www.yandex.ru",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config.ConfigApp.BaseURL = "http://example.com"
			// var router chi.Router
			router := NewRouter()

			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.originalURL))
			request.Header.Set("Content-Type", "text/plain")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, request)
			res := w.Result()
			defer func() {
				if err := res.Body.Close(); err != nil {
					log.Printf("close body: %v", err)
				}
			}()
			require.Equal(t, http.StatusCreated, res.StatusCode)
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			shortURL, err := url.Parse(string(resBody))
			require.NoError(t, err)

			gw := httptest.NewRecorder()
			getRequest := httptest.NewRequest(http.MethodGet, shortURL.Path, nil)
			getRequest.Header.Set("Content-Type", "text/plain")
			router.ServeHTTP(gw, getRequest)
			gres := gw.Result()
			defer func() {
				if err := gres.Body.Close(); err != nil {
					log.Printf("close body: %v", err)
				}
			}()
			assert.Equal(t, test.want.code, gres.StatusCode)
			assert.Equal(t, test.want.location, gres.Header.Get("Location"))
		})
	}
}
