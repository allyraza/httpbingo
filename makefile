default: config
	go build -o httpbin-server ./cmd/httpbin-server

linux: config
	CGO_ENABLED=0 GOOS=linux go build -o httpbin-server ./cmd/httpbin-server 


config:
	cp config.json.example config.json
