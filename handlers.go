package httpbin

import (
	"fmt"
	"net/http"
)

func (h *Httpbin) IP(w http.ResponseWriter, r *http.Request) {
	h.JSON(w, struct{ IP string }{IP: r.RemoteAddr})
}

func (h *Httpbin) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Httpbin) Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}
