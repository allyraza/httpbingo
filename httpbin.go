package httpbingo

import (
	"net/http"

	"github.com/allyraza/httpbingo/handler"
)

type HTTPBingo struct {
	Config  *Config
	Handler *http.ServeMux
}

func New(config *Config) *HTTPBingo {
	return &HTTPBingo{
		Config:  config,
		Handler: handler.New(),
	}
}
