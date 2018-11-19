package httpbin

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpbinNew(t *testing.T) {
	config := &Config{
		Addr: ":3000",
	}
	h := New(config)

	server := httptest.NewServer(h)

	defer server.Close()

	w, err := http.Get(server.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer w.Body.Close()

	got := w.StatusCode
	want := http.StatusOK

	if want != got {
		t.Errorf("Want %v, Got %v\n", want, got)
	}
}
