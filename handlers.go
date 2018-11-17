package httpbin

import (
	"fmt"
	"net/http"
)

func (h *Httpbin) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Httpbin) Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}
