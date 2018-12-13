package api

import "net/http"

// Routes init handler and registers routes
func Routes(mux *http.ServeMux) {
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/post", PostHandler)
	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/ip", IPHandler)
	mux.HandleFunc("/user-agent", UserAgentHandler)
	mux.HandleFunc("/headers", HeadersHandler)
	mux.HandleFunc("/cache", CacheHandler)
}
