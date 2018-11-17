package httpbin

import "net/http"

type Httpbin struct {
	config *Config
	mux    *http.ServeMux
}

func New(config *Config) *Httpbin {
	h := &Httpbin{
		config: config,
	}

	h.Routes()

	return h
}

func (h *Httpbin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
