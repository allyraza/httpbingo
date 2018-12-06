package assert

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// StatusOK asserts if response code is ok
func StatusOK(t *testing.T, w *httptest.ResponseRecorder) {
	Equal(t, w.Code, http.StatusOK)
}

// ContentType asserts is of type contentType
func ContentType(t *testing.T, w *httptest.ResponseRecorder, contentType string) {
	Equal(t, w.Header().Get("Content-Type"), contentType)
}

// ContentTypeJSON asserts is contentType is json
func ContentTypeJSON(t *testing.T, w *httptest.ResponseRecorder) {
	Equal(t, w.Header().Get("Content-Type"), "application/json")
}
