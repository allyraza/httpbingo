package httpbin

import "net/http"

func (h *Httpbin) Routes() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/version", h.Version)
	mux.HandleFunc("/ip", h.IP)
	mux.HandleFunc("/user-agent", h.UserAgent)

	h.mux = mux
}
