package model

import (
	"fmt"
	"net/http"
)

type Header struct {
	Headers map[string]string `json:"headers"`
}

func (h *Header) Parse(r *http.Request) {
	headers := make(map[string]string, len(r.Header))

	for k, v := range r.Header {
		headers[k] = v[0]
	}

	fmt.Printf("%v\n", headers)

	h.Headers = headers
}
