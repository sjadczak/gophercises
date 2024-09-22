package main

import (
	"flag"
	"fmt"
	"time"
)

type Game struct {
	CsvPath string
	Limit   time.Duration
	Quiz
}

type Quiz struct {
	Problems []Problem
	Score    uint
}

type Problem struct {
	Question string
	Answer   string
}

func main() {
	// Set up a new Game
	g := &Game{}

	// Parse command line arguments to create Game
	flag.StringVar(&g.CsvPath, "csv", "problems.csv", "A CSV file of questions and answers in the format `question,answer`.")
	flag.DurationVar(&g.Limit, "limit", 30*time.Second, "Time limit for the quiz")
	flag.Parse()

	fmt.Println("Welcome to Quizr!!!")
	fmt.Printf("%+v\n", g)
}
