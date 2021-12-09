package day5

import (
	"fmt"
	"strings"

	"github.com/cwood/aoc2021/internal/parse"
)

type Position string

const (
	Horizonal Position = "h"
	Vertical  Position = "v"
	Diagnol   Position = "d"
)

type Line struct {
	Start    []int
	End      []int
	Position Position
}

func (l *Line) Range() [][]int {
	xdiff := l.Start[0] - l.End[0]
	ydiff := l.Start[1] - l.End[0]

	var rng = make([][]int, 0)

	rng = append(rng, l.Start)

	for i := 0; i <= xdiff; i++ {

	}

	rng = append(rng, l.End)
	return rng
}

func (l *Line) String() string {
	return fmt.Sprintf(
		"%s: %d,%d -> %d,%d",
		l.Position,
		l.Start[0],
		l.Start[1],

		l.End[0],
		l.End[1],
	)
}

var NonPoint = []int{0, 0}

func (l *Line) Intersecting(i *Line) (bool, []int) {
	if i.Position == Diagnol || l.Position == Diagnol {
		return false, NonPoint
	}

	switch i.Position {
	case Horizonal:
	case Vertical:
	}

	return false, NonPoint
}

func New(line string) *Line {
	pos := strings.SplitN(line, "->", 3)

	start := pos[0]
	end := pos[1]

	startPos := parse.MustAsInt(parse.AsInt(strings.TrimSpace(start)))
	endPos := parse.MustAsInt(parse.AsInt(strings.TrimSpace(end)))

	var p Position

	if startPos[0] == endPos[0] {
		p = Horizonal
	} else if startPos[1] == endPos[1] {
		p = Vertical
	} else {
		p = Diagnol
	}

	return &Line{
		Start:    startPos,
		End:      endPos,
		Position: p,
	}

}

func LinesIntersect(vls []*Line) int {
	var intersects map[string]int

	for _, vl := range vls {
		for _, ol := range vls {
			if ok, i := vl.Intersecting(ol); ok {
				key := fmt.Sprintf("%d,%d", i[0], i[1])
				if _, ok := intersects[key]; !ok {
					intersects[fmt.Sprintf("%d,%d", i[0], i[1])] = 1
				} else {
					intersects[fmt.Sprintf("%d,%d", i[0], i[1])] += 1
				}
			}
		}
	}

	var lineCount int

	for _, num := range intersects {
		if num >= 2 {
			lineCount += 1
		}
	}

	return lineCount
}
