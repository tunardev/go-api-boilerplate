package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	defaultPort = 8080
)

// Config is the configuration for the interval service.
type Config struct {
	Port int `json:"port"`
}

// LoadConfig loads the configuration from the given file.
func Load(file string) (*Config, error) {
	c := Config{
		Port: defaultPort,
	}

	// load from JSON config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	return &c, nil
}