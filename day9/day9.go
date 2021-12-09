package day9

import (
	"fmt"
	"sort"
)

var positions = [][]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func genChecks(cp []int, maxRow, maxCol int) [][]int {
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

func checkSurround(points [][]int, pos []int, checked map[string]int) {
	for _, n := range genChecks(pos, len(points)-1, len(points[0])-1) {
		pv := points[n[0]][n[1]]
		if pv == 9 {
			continue
		}
		key := fmt.Sprintf("%d,%d", n[0], n[1])
		if _, ok := checked[key]; !ok {
			checked[key] = pv
			checkSurround(points, n, checked)
		}
	}
}

func BasinsInMap(points [][]int) int {
	var basins = make([]int, 0)
	for _, lp := range lowPointsInMap(points) {
		var checked = make(map[string]int)
		checkSurround(points, lp.Pos, checked)

		var basinCnt int

		for _, _ = range checked {
			basinCnt += 1
		}

		basins = append(basins, basinCnt)
	}

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

type position struct {
	Pos []int
	Val int
}

func lowPointsInMap(points [][]int) []position {
	var lowPoints = make([]position, 0)
	for r := 0; r <= len(points)-1; r++ {
	columnLoop:
		for c := 0; c <= len(points[r])-1; c++ {
			pv := points[r][c]
			if pv == 9 {
				continue
			}

			checks := genChecks([]int{r, c}, len(points)-1, len(points[r])-1)
			for _, c := range checks {
				spotNum := points[c[0]][c[1]]
				if spotNum < pv {
					continue columnLoop
				}
			}

			lowPoints = append(lowPoints, position{
				Pos: []int{r, c},
				Val: pv,
			})
		}
	}

	return lowPoints
}

func LowpointInMap(points [][]int) int {
	lowPoints := lowPointsInMap(points)

	var pointSumN1 int
	for _, p := range lowPoints {
		pointSumN1 += p.Val + 1
	}
	return pointSumN1
}
