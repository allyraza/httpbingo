package model

import "net/http"

type IP struct {
	IP string `json:"ip"`
}

func (ip *IP) Parse(r *http.Request) {
	ip.IP = r.RemoteAddr
}
