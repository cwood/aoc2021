package day9

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

func LowpointInMap(points [][]int) int {
	var lowPoints = make([]int, 0)
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
			lowPoints = append(lowPoints, pv)
		}
	}

	var pointSumN1 int
	for _, p := range lowPoints {
		pointSumN1 += p + 1
	}
	return pointSumN1
}
