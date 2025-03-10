package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	urlshort "github.com/sjadczak/gophercises/urlshort/pkg"
)

func main() {
	// Handle CLI flags
	var yaml_path string
	var json_path string
	flag.StringVar(&yaml_path, "yaml", "", "Path to YAML file containing path-to-urls data.")
	flag.StringVar(&json_path, "json", "", "Path to JSON file containing path-to-urls data.")
	flag.Parse()

	// Create fallback http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	// Create path-to-url map
	paths := map[string]string{
		"/g":    "https://google.com",
		"/dags": "https://tenor.com/view/dags-do-you-like-dags-snatch-brad-pitt-gif-16273104",
	}

	// Create MapHandler
	mh := urlshort.MapHandler(paths, mux)

	// Create YAMLHandler
	yh, err := urlshort.YAMLHandler(yaml_path, mh)
	if err != nil {
		log.Fatalf("Couldn't parse YAML: %v", err)
		os.Exit(1)
	}

	// Create JSONHandler
	jh, err := urlshort.JSONHandler(json_path, yh)
	if err != nil {
		log.Fatalf("Couldn't parse JSON: %v", err)
		os.Exit(1)
	}

	// Run server & listen
	log.Println("URLShort listening on localhost:8080...")
	log.Println("Enter ctrl+c to quit")
	err = http.ListenAndServe("localhost:8080", jh)
	if err != nil {
		log.Fatalf("Couldn't start urlshort server: %v", err)
		os.Exit(1)
	}
}
