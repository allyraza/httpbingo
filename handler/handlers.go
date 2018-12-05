package handler

import (
	"fmt"
	"net/http"

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

func IP(w http.ResponseWriter, r *http.Request) {
	data := struct {
		IP string `json:"origin"`
	}{
		IP: r.RemoteAddr,
	}

	render.JSON(w, data)
}

func UserAgent(w http.ResponseWriter, r *http.Request) {
	data := struct {
		UserAgent string `json:"user-agent"`
	}{
		UserAgent: r.UserAgent(),
	}

	render.JSON(w, data)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, version.Full)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, hometpl)
}
