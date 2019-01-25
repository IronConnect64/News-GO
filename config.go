package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Configuration holds our configuration data.
type Configuration struct {
	Address string `json:"address"`
}

// LoadConfig loads the configuration into memory.
func LoadConfig() Configuration {
	cfg := Configuration{}
	file, err := ioutil.ReadFile("data/config.json")
	if err != nil {
		log.Fatalf("Failed to read config: %s\n", err.Error())
	}
	err = json.Unmarshal([]byte(file), &cfg)
	if err != nil {
		log.Fatalf("Failed to decode config: %s\n", err.Error())
	}
	return cfg
}
