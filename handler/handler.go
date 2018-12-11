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
		<title>httpbingo - a http test service</title>
	</head>
	<body>
	 <h1>httpbingo</h1>
	 <a href="#">Documentation</a>
	</body>
</html>
`

// New init handler and registers routes
func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/post", Post)
	mux.HandleFunc("/health", Health)
	mux.HandleFunc("/version", Version)
	mux.HandleFunc("/ip", IP)
	mux.HandleFunc("/user-agent", UserAgent)
	mux.HandleFunc("/headers", Headers)

	return mux
}

// Headers Handler
func Headers(w http.ResponseWriter, r *http.Request) {
	header := &model.Header{}
	header.Parse(r)

	render.JSON(w, header)
}

// IP Handler
func IP(w http.ResponseWriter, r *http.Request) {
	ip := model.IP{}
	ip.Parse(r)

	render.JSON(w, ip)
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

// Post accepts posts requests only
func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	fmt.Fprintf(w, "hello, %s", r.FormValue("name"))
}
