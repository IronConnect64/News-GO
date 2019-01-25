package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("NEWS-GO - 1.0.0")
	log.Println("Loading configuration...")
	config := LoadConfig()
	log.Println("Starting server...")
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}
