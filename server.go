package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("News-GO - 1.0.0")
	log.Println("Loading configuration...")
	config := Load()
	log.Println("Starting server...")
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}
