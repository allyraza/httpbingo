package handler_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allyraza/httpbingo"
	"github.com/allyraza/httpbingo/assert"
)

func request(method string, url string) *httptest.ResponseRecorder {
	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()

	h := httpbin.New(&httpbin.Config{})
	h.Handler.ServeHTTP(w, r)

	return w
}

func TestHomepage(t *testing.T) {
	w := request("GET", "/")

	assert.StatusOK(t, w.Code)
	assert.Contains(t, w.Body.String(), "httpbin")
}
