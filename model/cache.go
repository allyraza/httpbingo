package model

import "net/url"

type Cache struct {
	*IP
	*Header
	Query url.Values `json:"query"`
	URL   string     `json:"url"`
}
