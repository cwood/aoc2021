package day11

import (
	"github.com/cwood/aoc2021/internal/grid"
)

func flashOctopi(f [][]int, flashes *grid.IntSet, cpos []int) {
	tpos := grid.GenPositionsAll(cpos, len(f)-1, len(f[0])-1)
	for _, pos := range tpos {
		v := f[pos[0]][pos[1]]
		if v >= 9 {
			f[pos[0]][pos[1]] = 0
			flashes.Add(pos)
			flashOctopi(f, flashes, pos)
			continue
		}

		if v == 0 && flashes.Contains(pos) {
			continue
		}

		f[pos[0]][pos[1]]++
	}
}

func genOctopi(octopi [][]int) ([][]int, int) {
	var flashes = grid.NewIntSet()

	for r := 0; r <= len(octopi)-1; r++ {
		for c := 0; c <= len(octopi[r])-1; c++ {
			octi := octopi[r][c]
			if octi >= 9 {
				octopi[r][c] = 0
				flashes.Add([]int{r, c})
				flashOctopi(octopi, flashes, []int{r, c})
				octi = 0
			} else {
				if flashes.Contains([]int{r, c}) {
					continue
				}

				octopi[r][c]++
			}

		}
	}

	return octopi, flashes.Count()
}

func OctopusFlashesPerTurn(octopi [][]int, turns int) int {

	var flashes int

	for t := 0; t <= turns-1; t++ {
		var f int
		octopi, f = genOctopi(octopi)
		flashes += f
	}

	return flashes
}
