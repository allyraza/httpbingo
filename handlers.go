package httpbin

import (
	"fmt"
	"net/http"
)

func (h *Httpbin) IP(w http.ResponseWriter, r *http.Request) {
	h.JSON(w, struct {
		IP string `json:"origin"`
	}{IP: r.RemoteAddr})
}

func (h *Httpbin) UserAgent(w http.ResponseWriter, r *http.Request) {
	h.JSON(w, struct {
		UserAgent string `json:"user-agent"`
	}{UserAgent: r.UserAgent()})
}

func (h *Httpbin) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Httpbin) Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version)
}
