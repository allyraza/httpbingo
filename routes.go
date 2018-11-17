package httpbin

import "net/http"

func (h *Httpbin) Routes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/version", h.Version)

	h.mux = mux
}
