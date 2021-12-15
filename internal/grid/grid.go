package grid

import (
	"fmt"
	"log"
)

var positions = [][]int{
	{-1, 0},  // Up
	{-1, 1},  // Right and up
	{0, 1},   // Right
	{1, 0},   // Down
	{1, 1},   // Down and right
	{1, -1},  // Left and down
	{0, -1},  // Left
	{-1, -1}, // Left and up
}

var direct = [][]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func genPositions(positions [][]int, cp []int, mr, mc int) [][]int {
	var poss = make([][]int, 0)

	for _, p := range positions {
		gened := []int{
			p[0] + cp[0],
			p[1] + cp[1],
		}

		if gened[0] < 0 || gened[1] < 0 {
			continue
		} else if gened[0] > mr || gened[1] > mc {
			continue
		}

		poss = append(poss, gened)
	}

	return poss
}

func GenPositionDirect(cp []int, maxRow, maxCol int) [][]int {
	return genPositions(direct, cp, maxRow, maxCol)
}

func GenPositionsAll(cp []int, maxRow, maxCol int) [][]int {
	return genPositions(positions, cp, maxRow, maxCol)
}

func Printf(g [][]int, msg string, args ...interface{}) {
	m := "\n"
	for r := 0; r <= len(g)-1; r++ {
		for c := 0; c <= len(g[r])-1; c++ {
			m += fmt.Sprintf("| %d ", g[r][c])
		}
		m += "\n"
	}
	log.Printf("%s %s\n", fmt.Sprintf(msg, args...), m)
}

type IntSet struct {
	i map[string]bool
}

func (i *IntSet) Add(p []int) {
	key := fmt.Sprintf("%d.%d", p[0], p[1])
	if _, ok := i.i[key]; !ok {
		i.i[key] = true
	}
}

func (i *IntSet) Contains(p []int) bool {
	key := fmt.Sprintf("%d.%d", p[0], p[1])
	_, ok := i.i[key]
	return ok
}

func (i *IntSet) Count() int {
	return len(i.i)
}

func Len(g [][]int) int {
	return len(g) * len(g[0])
}

func NewIntSet() *IntSet {
	return &IntSet{i: make(map[string]bool)}
}
