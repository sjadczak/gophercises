package quizr

import (
	"cmp"
	"encoding/csv"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	CURLINE = "\x1b[0G\x1b[2K"
	PRVLINE = "\x1b[1F\x1b[2K"
)

const (
	Done gameState = iota
	Timeout
)

type gameState int

func (gs gameState) escapeCode() string {
	if gs == Done {
		return PRVLINE
	}

	return CURLINE
}

type Game struct {
	// Game configuration
	CsvPath string
	Limit   time.Duration
	Shuffle bool

	// Internal game state
	score    uint
	problems []Problem
	inputs   []string
	endState gameState
}

func (g *Game) Init() error {
	fd, err := os.Open(g.CsvPath)
	if err != nil {
		msg := fmt.Sprintf("Couldn't open `%s`, is this a valid filepath?", g.CsvPath)
		return PubError(ErrCSVPath, msg)
	}
	defer fd.Close()

	// Read in the CSV file
	r := csv.NewReader(fd)
	ps, err := r.ReadAll()
	if err != nil {
		return PubError(ErrNotCSV, "This doesn't seem to be a csv file...")
	}

	// Parse the csv records into the Quiz struct, returned error is guaranteed to be PublicError
	err = g.parseQuestions(ps)
	if err != nil {
		return err
	}

	// If g.Shuffle, shuffle the question order
	if g.Shuffle {
		g.shuffleQuestions()
	}

	return nil
}

func (g *Game) shuffleQuestions() {
	rand.Shuffle(len(g.problems), func(i, j int) {
		g.problems[i], g.problems[j] = g.problems[j], g.problems[i]
	})
}

func (g *Game) parseQuestions(problems [][]string) error {
	for _, problem := range problems {
		if len(problem) != 2 {
			msg := fmt.Sprintf("This csv files seems to be formatted incorrectly. There should be two columns, not %d", len(problem))
			return PubError(ErrCSVFormat, msg)
		}

		p := Problem{
			q: problem[0],
			a: strings.TrimSpace(problem[1]),
		}

		g.problems = append(g.problems, p)
	}

	return nil
}

func (g *Game) runQuiz(answers chan<- bool) {
	defer close(answers)
	pad := padding(len(g.problems))

	for i, problem := range g.problems {
		problem.askQuestion(i, pad)
		res, ans := problem.checkResponse()
		g.inputs = append(g.inputs, ans)
		answers <- res
	}
}

func (g *Game) printResults() {
	fmt.Print(g.endState.escapeCode())

	if g.endState == Timeout {
		fmt.Println("You ran out of time...")
		fmt.Printf("You only completed %d of %d questions, better luck next time!\n", len(g.inputs), len(g.problems))
	}

	a_max := slices.MaxFunc(g.problems, func(a, b Problem) int {
		ai := a.padding("a")
		bi := b.padding("a")

		return cmp.Compare(ai, bi)
	})

	i_max := slices.MaxFunc(g.inputs, func(a, b string) int {
		ai, _ := strconv.Atoi(a)
		bi, _ := strconv.Atoi(b)

		return cmp.Compare(ai, bi)
	})

	q_pad := padding(len(g.problems))
	a_pad := a_max.padding("a")

	ii, _ := strconv.Atoi(i_max)
	i_pad := padding(ii)

	for i, in := range g.inputs {
		p := g.problems[i]

		var comp string
		if strings.EqualFold(p.a, strings.TrimSpace(in)) {
			comp = "=="
		} else {
			comp = "!="
		}

		fmt.Printf("Q%0*d) %s: %*s %s %*s\n", q_pad, i+1, p.q, a_pad, p.a, comp, i_pad, in)
	}
}

func (g *Game) Run() {
	fmt.Println("Welcome to Quizr!")

	timer := time.NewTimer(g.Limit)
	results := make(chan bool)

	// Start goroutine to ask questions and return answers of channel
	go g.runQuiz(results)

	// Run game loop
	for {
		select {
		case <-timer.C:
			g.endState = Timeout
			return
		case correct, ok := <-results:
			if correct {
				g.score++
			}

			if !ok {
				timer.Stop()
				g.endState = Done
				return
			}
		}
	}
}

func (g *Game) Finalize() {
	g.printResults()
	fmt.Printf("You answered %d of %d correctly!\n", g.score, len(g.problems))
}

// Determines the number of decimal digits in an integer to set padding
// on display
func padding(x int) int {
	if x == 0 {
		return 1
	}

	// while x isn't 0, integer divide by 10 and increment count
	count := 0
	for x != 0 {
		x /= 10
		count++
	}

	return count
}
