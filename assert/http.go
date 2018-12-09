package assert

import (
	"net/http/httptest"
	"strings"
	"testing"
)

// Status asserts if response code is ok
func Status(t *testing.T, w *httptest.ResponseRecorder, statusCode int) {
	Equal(t, w.Code, statusCode)
}

// Header asserts if header has value
func Header(t *testing.T, w *httptest.ResponseRecorder, name string, value interface{}) {
	Equal(t, w.Header().Get(name), value)
}

// Body asserts is equal to body
func Body(t *testing.T, w *httptest.ResponseRecorder, want string) {
	body := strings.TrimSpace(w.Body.String())
	Equal(t, want, body)
}

// BodyContains asserts if text is in body
func BodyContains(t *testing.T, w *httptest.ResponseRecorder, want string) {
	Contains(t, w.Body.String(), want)
}

// ContentType asserts is of type contentType
func ContentType(t *testing.T, w *httptest.ResponseRecorder, contentType string) {
	Header(t, w, "Content-Type", contentType)
}
