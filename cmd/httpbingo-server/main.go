package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/allyraza/httpbingo"
)

func main() {
	config := &httpbingo.Config{}

	flag.StringVar(&config.Filepath, "config", "config.json", "Config file for httpbingo-server.")
	flag.Parse()

	config.ParseFile()

	log.Printf("Starting server on %v\n", config.Address)
	app := httpbingo.New(config)

	server := http.Server{
		Addr:    config.Address,
		Handler: app.Handler,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	log.Printf("Shutdown signal received, quiting...")
	server.Shutdown(context.Background())
}
