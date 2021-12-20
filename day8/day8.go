package day8

import "github.com/cwood/aoc2021/internal/parse"

type Segment int

const (
	One   Segment = 2
	Four  Segment = 4
	Seven Segment = 3
	Eight Segment = 7
)

type Light struct {
	Top         string
	RightTop    string
	LeftTop     string
	Middle      string
	RightBottom string
	LeftBottom  string
	Bottom      string
}

func setIfEmpty(s *string, val byte) {
	if *s == "" {
		*s = string(val)
	}
}

func (l *Light) ToNum(a string) int {
	switch Segment(len(a)) {
	case One:
		setIfEmpty(&l.RightTop, a[0])
		setIfEmpty(&l.RightBottom, a[1])
		return 1
	case Four:
		return 4
	case Seven:
		return 7
	case Eight:
		return 7
	}
	return 0
}

func CountOutputSums(inout []parse.InOut) int {
	return 0
}

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
