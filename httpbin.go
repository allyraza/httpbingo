package httpbin

import (
	"net/http"

	"github.com/allyraza/httpbingo/handler"
)

type HTTPBin struct {
	config *Config
	mux    *http.ServeMux
}

func New(config *Config) *HTTPBin {
	h := &HTTPBin{
		config: config,
	}

	h.mux = h.NewMux()

	return h
}

func (h *HTTPBin) NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/version", handler.Version)
	mux.HandleFunc("/ip", handler.IP)
	mux.HandleFunc("/user-agent", handler.UserAgent)

	return mux
}

func (h *HTTPBin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
