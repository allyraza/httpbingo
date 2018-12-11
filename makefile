default: config
	go build -o httpbingo-server ./cmd/httpbingo-server

linux: config
	CGO_ENABLED=0 GOOS=linux go build -o httpbingo-server ./cmd/httpbingo-server 


config:
	cp config.json.example config.json
