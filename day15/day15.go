package day15

import "sort"

type pos struct {
	p []int
	v int
}

func LowestTotalRisk(g [][]int) int {
	var risks []int

	for {
		cpos := []int{0, 0}
		tpos := []int{len(g), len(g[0]) - 1}

		crisk := make([]pos, 0)

		for _, pos := grid.GenPositionDirect(cpos, len(g), len(g[0])) {
		}

	}

	sort.Ints(risks)
	return risks[len(risks)-1]
}
