package day6

func LanternFishAfterNDays(initial []int, n int) int {

	var fish = make(map[int]int)

	for _, i := range initial {
		if _, ok := fish[i]; !ok {
			fish[i] = 1
		} else {
			fish[i]++
		}
	}

	for i := n; i > 0; i-- {
		fishDay := make(map[int]int)

		for j := 8; j >= 0; j-- {
			if n, ok := fish[j]; ok {
				if j != 0 {
					fishDay[j-1] = n
				}
				if j == 0 {
					fishDay[8] += n
					fishDay[6] += n
				}
			}
		}

		fish = fishDay
	}

	var sum int

	for _, n := range fish {
		sum += n
	}

	return sum
}
