package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Content contains all Content objects.
type Content struct {
	Contents []ContentObject `json:"content"`
}

// ContentObject represents RSS content.
type ContentObject struct {
	DataPath string `json:"datapath"`
	Name     string `json:"name"`
	Output   string `json:"output"`
	URL      string `json:"url"`
}

// LoadContent loads the contents into memory.
func LoadContent() Content {
	content := Content{}
	file, err := ioutil.ReadFile("data/content.json")
	if err != nil {
		log.Fatalf("Failed to read content: %s\n", err.Error())
	}
	err = json.Unmarshal([]byte(file), &content)
	if err != nil {
		log.Fatalf("Failed to decode content: %s\n", err.Error())
	}
	return content
}
