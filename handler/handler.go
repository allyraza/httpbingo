package handler

import (
	"fmt"
	"net/http"

	"github.com/allyraza/httpbingo/model"
	"github.com/allyraza/httpbingo/render"
	"github.com/allyraza/httpbingo/version"
)

const hometpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<title>httpbin - a http test service</title>
	</head>
	<body>
	 <h1>httpbin</h1>
	 <a href="#">Documentation</a>
	</body>
</html>
`

// New init handler and registers routes
func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/version", Version)
	mux.HandleFunc("/ip", IP)
	mux.HandleFunc("/user-agent", UserAgent)

	return mux
}

// IP Handler
func IP(w http.ResponseWriter, r *http.Request) {
	data := model.IP{
		IP: r.RemoteAddr,
	}

	render.JSON(w, data)
}

// UserAgent Handler
func UserAgent(w http.ResponseWriter, r *http.Request) {
	data := model.UserAgent{
		UserAgent: r.UserAgent(),
	}

	render.JSON(w, data)
}

// Health Handler
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Version Handler
func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version.Full)
}

// Home Handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, hometpl)
}
