package httpbin

import (
	"net/http"

	"github.com/allyraza/httpbingo/handler"
)

type HTTPBin struct {
	Config  *Config
	Handler *http.ServeMux
}

func New(config *Config) *HTTPBin {
	return &HTTPBin{
		Config:  config,
		Handler: handler.New(),
	}
}
