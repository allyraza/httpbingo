package request

import (
	"net/http"
)

// ParseIP parses and returns request address.
func ParseIP(r *http.Request) string {
	return r.RemoteAddr
}

// ParseHeaders returns headers a map
func ParseHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string, len(r.Header))

	for k, v := range r.Header {
		headers[k] = v[0]
	}

	return headers
}
