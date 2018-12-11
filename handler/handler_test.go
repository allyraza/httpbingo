package handler_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/allyraza/httpbingo"
	"github.com/allyraza/httpbingo/assert"
)

type RequestFunc func(string, url.Values) *httptest.ResponseRecorder

func MakeRequest(method string) RequestFunc {
	return func(url string, params url.Values) *httptest.ResponseRecorder {
		r, err := http.NewRequest(method, url, strings.NewReader(params.Encode()))
		if err != nil {
			log.Fatal(err)
		}

		r.RemoteAddr = "127.0.0.1"
		r.Header.Set("user-agent", "HTTPBingo")
		r.Header.Set("Accept", "application/json; charset=utf-8")
		r.Header.Set("Host", "127.0.0.1")

		if method == http.MethodPost {
			r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		}

		w := httptest.NewRecorder()

		h := httpbingo.New(&httpbingo.Config{})
		h.Handler.ServeHTTP(w, r)

		return w
	}
}

func Get(url string, params url.Values) *httptest.ResponseRecorder {
	return MakeRequest(http.MethodGet)(url, params)
}

func Post(url string, params url.Values) *httptest.ResponseRecorder {
	return MakeRequest(http.MethodPost)(url, params)
}

func TestHomepage(t *testing.T) {
	w := Get("/", url.Values{})

	assert.StatusOK(t, w)
	assert.BodyContains(t, w, "httpbingo")
}

func TestIP(t *testing.T) {
	w := Get("/ip", url.Values{})

	assert.StatusOK(t, w)
	assert.JSON(t, w, `{"ip":"127.0.0.1"}`)
}

func TestHeaders(t *testing.T) {
	w := Get("/headers", url.Values{})

	headers := struct {
		Headers map[string]string `json:"headers"`
	}{
		Headers: map[string]string{
			"Accept":     "application/json; charset=utf-8",
			"User-Agent": "HTTPBingo",
			"Host":       "127.0.0.1",
		},
	}

	want, _ := json.Marshal(headers)

	assert.StatusOK(t, w)
	assert.JSON(t, w, string(want))
}

func TestUserAgent(t *testing.T) {
	w := Get("/user-agent", url.Values{})

	assert.StatusOK(t, w)
	assert.JSON(t, w, `{"user-agent":"HTTPBingo"}`)
}

func TestPost(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		w := Post("/post", url.Values{"name": []string{"foo"}})

		assert.StatusOK(t, w)
		assert.Body(t, w, "hello, foo")
	})

	t.Run("Error", func(t *testing.T) {
		w := Get("/post", url.Values{"name": []string{"foo"}})

		assert.Status(t, w, http.StatusMethodNotAllowed)
		assert.Body(t, w, "Method not allowed")
	})
}
