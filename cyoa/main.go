package main

import (
	"fmt"
	"log"

	"github.com/sjadczak/gophercises/cyoa/story"
)

func main() {
	s, err := story.ParseStory("gopher.json")
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Debug printing story:")
	fmt.Printf("%+v\n", s)
}
