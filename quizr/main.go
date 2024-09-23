package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sjadczak/gophercises/quizr/quizr"
)

func main() {
	// Set up a new Game
	g := &quizr.Game{}

	// Parse command line arguments to create Game
	flag.StringVar(&g.CsvPath, "csv", "problems.csv", "A CSV file of questions and answers in the format `question,answer`.")
	flag.DurationVar(&g.Limit, "limit", 30*time.Second, "Time limit for the quiz")
	flag.BoolVar(&g.Shuffle, "shuffle", false, "Shuffle the questions before starting?")
	flag.Parse()

	// Initialize Game state
	err := g.Init()
	if err != nil {
		var pubError *quizr.PublicError
		if quizr.As(err, &pubError) {
			fmt.Println(pubError.Public())
			os.Exit(1)
		} else {
			panic(err)
		}
	}

	// Run game, present final results
	g.Run()
	g.Finalize()
}
