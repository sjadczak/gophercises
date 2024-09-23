package quizr

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem struct {
	q string
	a string
}

func (p Problem) askQuestion(i, pad int) {
	esc := ""
	if i != 0 {
		esc = PRVLINE
	}

	fmt.Printf("%sQ%0*d) %s: ", esc, pad, i+1, p.q)
}

func (p Problem) checkResponse() (bool, string) {
	var ans string
	fmt.Scanf("%s\n", &ans)

	res := strings.EqualFold(p.a, strings.TrimSpace(ans))
	return res, ans
}

func (p Problem) padding(field string) int {
	pad := 1

	if field == "a" {
		ai, _ := strconv.Atoi(p.a)
		pad = padding(ai)
	} else if field == "q" {
		pad = len(p.q)
	}

	return pad
}
