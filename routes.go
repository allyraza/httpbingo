package httpbin

import "net/http"

func (h *Httpbin) Routes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/version", h.Version)
	mux.HandleFunc("/ip", h.IP)

	h.mux = mux
}
