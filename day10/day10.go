package day10

import (
	"fmt"
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

func isLineCorrupt(l string) (bool, string) {
	var toFind string
	var tokens = strings.Split(l, "")

tokenScan:
	for tp, t := range tokens {
		switch t {
		case "(":
			toFind = ")"
		case "[":
			toFind = "]"
		case "<":
			toFind = ">"
		case "{":
			toFind = "}"
		}

		for i := len(tokens) - 1; i >= tp; i-- {
			nt := tokens[i]
			if nt == toFind {
				continue tokenScan
			}
		}
	}
	return false, noCorrupts
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
		case ")":
			score += n * ParenScore
		case "]":
			score += n * BlockScore
		case "}":
			score += n * CurvyScore
		case ">":
			score += n * SignScore
		default:
			panic(fmt.Sprintf("unknown sign: %s", p))
		}
	}

	return 0
}
