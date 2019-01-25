package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Content represents RSS content.
type Content struct {
	DataPath string `json:"datapath"`
	Name     string `json:"name"`
	Output   string `json:"output"`
	URL      string `json:"url"`
}

// Configuration holds our configuration data.
type Configuration struct {
	Address  string    `json:"address"`
	Contents []Content `json:"content"`
}

// Load loads the properties file into memory.
func Load() Configuration {
	cfg := Configuration{}
	file, err := ioutil.ReadFile("properties.json")
	if err != nil {
		log.Fatalf("Failed to read properties: %s\n", err.Error())
	}
	err = json.Unmarshal([]byte(file), &cfg)
	if err != nil {
		log.Fatalf("Failed to decode properties: %s\n", err.Error())
	}
	return cfg
}
