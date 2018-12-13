package api

import "net/url"

type Cache struct {
	Query   url.Values        `json:"query"`
	Headers map[string]string `json:"headers"`
	IP      string            `json:"ip"`
	URL     string            `json:"url"`
}
