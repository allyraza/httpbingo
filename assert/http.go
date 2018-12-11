package assert

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Status asserts if response code is equal to statusCode
func Status(t *testing.T, w *httptest.ResponseRecorder, statusCode int) {
	Equal(t, statusCode, w.Code)
}

// StatusOK asserts if response code is ok
func StatusOK(t *testing.T, w *httptest.ResponseRecorder) {
	Equal(t, w.Code, http.StatusOK)
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

// JSON asserts validates json header and body
func JSON(t *testing.T, w *httptest.ResponseRecorder, body string) {
	Header(t, w, "Content-Type", "application/json; charset=utf-8")
	Body(t, w, body)
}
