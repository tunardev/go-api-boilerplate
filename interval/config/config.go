package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	defaultPort = 8080
	Mongo_URI = "mongodb://localhost:27017"
)

// Config is the configuration for the app service.
type Config struct {
	Port int `json:"port"`
	Mongo_URI string `json:"mongodb"`
}

// Load loads the configuration from the given file.
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