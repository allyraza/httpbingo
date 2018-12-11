package httpbingo

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config parses and holds config for httpbingo server.
type Config struct {
	Filepath string
	Address  string `json:"address"`
}

// ParseFile parses config file from given file path.
func (c *Config) ParseFile() {
	blob, err := ioutil.ReadFile(c.Filepath)
	if err != nil {
		log.Fatalf("Config file: %v\n", err)
	}

	if err := json.Unmarshal(blob, c); err != nil {
		log.Fatalf("Config parse: %v\n", err)
	}
}
