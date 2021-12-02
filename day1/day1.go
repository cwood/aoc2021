package day1

func createWindows(ints []int, size int) []int {
	var (
		windowInts = make([]int, 0)
		intsLen    = len(ints) - 1
	)

createWindows:
	for i, _ := range ints {
		var windowSize = 0
		for j := 0; j <= size-1; j++ {
			index := i + j
			if index > intsLen {
				continue createWindows
			}
			windowSize += ints[index]
		}
		windowInts = append(windowInts, windowSize)
	}

	return windowInts
}

func CalculateIncreases(ints []int, windowSize int) int {
	if windowSize > 1 {
		ints = createWindows(ints, windowSize)
	}

	var increases = 0
	for i, n := range ints {
		if i > 0 {
			prev := ints[i-1]
			if prev < n {
				increases++
			}
		}
	}
	return increases
}
