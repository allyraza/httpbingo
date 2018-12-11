package handler_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allyraza/httpbingo"
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

	h := httpbingo.New(&httpbingo.Config{})
	h.Handler.ServeHTTP(w, r)

	return w
}

func TestHomepage(t *testing.T) {
	w := Get("/")

	assert.StatusOK(t, w)
	assert.BodyContains(t, w, "httpbingo")
}

func TestIP(t *testing.T) {
	w := Get("/ip")

	assert.StatusOK(t, w)
	assert.JSON(t, w, `{"origin":"127.0.0.1"}`)
}

func TestUserAgent(t *testing.T) {
	w := Get("/user-agent")

	assert.StatusOK(t, w)
	assert.JSON(t, w, `{"user-agent":"TestBrowser"}`)
}
