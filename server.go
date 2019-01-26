package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// ErrorCheck makes error checking easier.
func ErrorCheck(err error) {
	if err != nil {
		log.Fatalf("An error occurred: %s", err.Error())
	}
}

func main() {
	log.Println("News-GO - 1.0.0")

	log.Println("Loading configuration...")
	cfg := Config{}
	file, err := ioutil.ReadFile("config.json")
	ErrorCheck(err)
	ErrorCheck(json.Unmarshal([]byte(file), &cfg))

	// TODO: implementing RSS parsing

	log.Println("Starting server...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(cfg.Address, nil))

	// From there on, everything server-related will be redirected to our handler function.
}

func handler(w http.ResponseWriter, r *http.Request) {

}
