package httpbingo

import (
	"net/http"

	"github.com/allyraza/httpbingo/api"
)

type HTTPBingo struct {
	Config  *Config
	Handler *http.ServeMux
}

func New(config *Config) *HTTPBingo {
	mux := http.NewServeMux()
	api.Routes(mux)

	return &HTTPBingo{
		Config:  config,
		Handler: mux,
	}
}
