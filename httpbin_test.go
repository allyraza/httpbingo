package httpbingo_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allyraza/httpbingo"
	"github.com/allyraza/httpbingo/assert"
)

func TestHTTPBingoNew(t *testing.T) {
	config := &httpbingo.Config{
		Address: ":3000",
	}
	h := httpbingo.New(config)

	server := httptest.NewServer(h.Handler)

	defer server.Close()

	w, err := http.Get(server.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer w.Body.Close()

	assert.Equal(t, w.StatusCode, http.StatusOK)
}
