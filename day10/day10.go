package day10

import (
	"fmt"
	"sort"
	"strings"
)

const (
	ParenScore = 3
	BlockScore = 57
	CurvyScore = 1197
	SignScore  = 25137
)

var (
	noCorrupts = ""
)

func pop(s []string) (string, []string) {
	l := s[len(s)-1]
	return l, s[:len(s)-1]
}

func isLineCorrupt(l string) (bool, string) {
	var tokens = strings.Split(l, "")

	var openers = make([]string, 0)

	for _, t := range tokens {
		switch t {
		case ")":
			l, openers = pop(openers)
			if l != "(" {
				return true, "("
			}
		case "]":
			l, openers = pop(openers)
			if l != "[" {
				return true, "["
			}
		case ">":
			l, openers = pop(openers)
			if l != "<" {
				return true, "<"
			}
		case "}":
			l, openers = pop(openers)
			if l != "{" {
				return true, "{"
			}
		default:
			openers = append(openers, t)
		}
	}

	return false, strings.Join(openers, "")
}

func SyntaxScoringComplete(lines []string) int {

	var linePoints = make([]int, 0)

	for _, l := range lines {
		isCorrupt, c := isLineCorrupt(l)
		if isCorrupt {
			continue
		}

		var points = 0

		cs := strings.Split(c, "")
		for i := len(cs) - 1; i >= 0; i-- {
			points = points * 5
			ch := cs[i]
			switch ch {
			case "[":
				points += 2
			case "(":
				points += 1
			case "{":
				points += 3
			case "<":
				points += 4
			}
		}
		linePoints = append(linePoints, points)
	}

	sort.Ints(linePoints)

	return linePoints[int(len(linePoints)/2)]

}

func SyntaxScoringCorrupt(lines []string) int {
	var missingParts = make(map[string]int)

	for _, l := range lines {
		if isCorrupt, c := isLineCorrupt(l); isCorrupt {
			if _, ok := missingParts[c]; !ok {
				missingParts[c] = 1
			} else {
				missingParts[c]++
			}
		}
	}

	var score int
	for p, n := range missingParts {
		switch p {
		case "(":
			score += n * ParenScore
		case "[":
			score += n * BlockScore
		case "{":
			score += n * CurvyScore
		case "<":
			score += n * SignScore
		default:
			panic(fmt.Sprintf("unknown sign: %s", p))
		}
	}

	return score
}
