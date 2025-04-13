package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sjadczak/gophercises/cyoa/story"
	"github.com/sjadczak/gophercises/cyoa/views"
)

func main() {
	s, err := story.NewStoryService("gopher.json")
	if err != nil {
		log.Fatalf("%v", err)
	}
	s.Tpl, err = views.Parse("templates/story.gohtml")
	if err != nil {
		log.Fatalf("parsing template: %v", err)
	}

	http.HandleFunc("/", redirect)
	http.HandleFunc("/story/", s.WebHandler)

	fmt.Println("Starting CYOA Server on :8080...")
	http.ListenAndServe(":8080", nil)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/story/intro", http.StatusFound)
}
