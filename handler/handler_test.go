package handler_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	httpbin "github.com/allyraza/httpbingo"
	"github.com/allyraza/httpbingo/assert"
)

func Get(url string) *httptest.ResponseRecorder {
	r, err := http.NewRequest(http.MethodGet, url, nil)
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
	w := Get("/")

	assert.Status(t, w, http.StatusOK)
	assert.BodyContains(t, w, "httpbin")
}

func TestIP(t *testing.T) {
	w := Get("/ip")

	assert.Status(t, w, http.StatusOK)
	assert.ContentType(t, w, "application/json")
	assert.Body(t, w, `{"origin":"127.0.0.1"}`)
}

func TestUserAgent(t *testing.T) {
	w := Get("/user-agent")

	assert.Status(t, w, http.StatusOK)
	assert.ContentType(t, w, "application/json")
	assert.Body(t, w, `{"user-agent":"TestBrowser"}`)
}
