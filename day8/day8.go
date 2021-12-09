package day8

import "github.com/cwood/aoc2021/internal/parse"

type Segment int

const (
	One   Segment = 2
	Four  Segment = 4
	Seven Segment = 3
	Eight Segment = 7
)

func CountUniqueOutput(inout []parse.InOut) int {
	var uniqueSeg int
	for _, io := range inout {
		for _, out := range io.Out {
			switch Segment(len(out)) {
			case One,
				Four,
				Seven,
				Eight:
				uniqueSeg += 1
			}
		}
	}

	return uniqueSeg
}
