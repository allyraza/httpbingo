package api

import (
	"fmt"
	"net/http"

	"github.com/allyraza/httpbingo/render"
	"github.com/allyraza/httpbingo/request"
	"github.com/allyraza/httpbingo/version"
)

const hometpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<title>httpbingo - a http test service</title>
	</head>
	<body>
	 <h1>httpbingo</h1>
	 <a href="#">Documentation</a>
	</body>
</html>
`

// HeadersHandler
func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	header := &Header{
		Headers: request.ParseHeaders(r),
	}

	render.JSON(w, header)
}

// IPHandler
func IPHandler(w http.ResponseWriter, r *http.Request) {
	ip := IP{
		IP: request.ParseIP(r),
	}

	render.JSON(w, ip)
}

// UserAgentHandler
func UserAgentHandler(w http.ResponseWriter, r *http.Request) {
	data := UserAgent{
		UserAgent: r.UserAgent(),
	}

	render.JSON(w, data)
}

// CacheHandler
func CacheHandler(w http.ResponseWriter, r *http.Request) {
	cache := Cache{
		Query:   r.URL.Query(),
		Headers: request.ParseHeaders(r),
		IP:      request.ParseIP(r),
		URL:     r.URL.String(),
	}

	render.JSON(w, cache)
}

// Health Handler
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Version Handler
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version.Full)
}

// Home Handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, hometpl)
}

// Post accepts posts requests only
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	fmt.Fprintf(w, "hello, %s", r.FormValue("name"))
}
