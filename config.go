package httpbin

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config parses and holds config for httpbin server
type Config struct {
	Filepath string
	Addr     string `json:"addr"`
}

// ParseFile parses config file for given file path
func (c *Config) ParseFile() {
	blob, err := ioutil.ReadFile(c.Filepath)
	if err != nil {
		log.Fatalf("Config file: %v\n", err)
	}

	if err := json.Unmarshal(blob, c); err != nil {
		log.Fatalf("Config parse: %v\n", err)
	}
}
