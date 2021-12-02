package day1

func CalculateIncreases(ints []int) int {
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
