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

func GenPositionsAll(cp []int, maxRow, maxCol int) [][]int {
	var poss = make([][]int, 0)

	for _, p := range positions {
		gened := []int{
			p[0] + cp[0],
			p[1] + cp[1],
		}

		if gened[0] < 0 || gened[1] < 0 {
			continue
		} else if gened[0] > maxRow || gened[1] > maxCol {
			continue
		}

		poss = append(poss, gened)
	}

	return poss

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
