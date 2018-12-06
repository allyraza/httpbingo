package handler_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	httpbin "github.com/allyraza/httpbingo"
	"github.com/allyraza/httpbingo/assert"
)

func request(method string, url string) *httptest.ResponseRecorder {
	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	r.RemoteAddr = "127.0.0.1"
	r.Header.Set("user-agent", "TestBrowser")

	w := httptest.NewRecorder()

	h := httpbin.New(&httpbin.Config{})
	h.Handler.ServeHTTP(w, r)

	return w
}

func TestHomepage(t *testing.T) {
	w := request("GET", "/")

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(), "httpbin")
}

func TestIP(t *testing.T) {
	w := request("GET", "/ip")

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json")
	assert.Contains(t, w.Body.String(), `{"origin":"127.0.0.1"}`)
}

func TestUserAgent(t *testing.T) {
	w := request("GET", "/user-agent")

	assert.StatusOK(t, w)
	assert.ContentTypeJSON(t, w)
	assert.Contains(t, w.Body.String(), `{"user-agent":"TestBrowser"}`)
}
