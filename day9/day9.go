package day9

var positions = [][]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func genChecks(cp []int) [][]int {
	for _, p := range positions {

	}
}

func LowpointInMap(points [][]int) int {
	for r := 0; r <= len(points)-1; r++ {
		for c := 0; c <= len(points[r])-1; c++ {
			pv := points[r][c]
			if pv == 9 {
				continue
			}

		}
	}
	return 0

}
