package day7

import (
	"math"
	"sort"
)

func LeastFuelUsed(crabs []int, isConstant bool) int {
	sort.Ints(crabs)

	var leastFuelUsed int

	maxSize := crabs[len(crabs)-1]

	for i := 0; i <= maxSize; i++ {

		var fuelUsed int
		for j := len(crabs) - 1; j >= 0; j-- {
			fuelDelta := math.Abs(float64(i - crabs[j]))
			if !isConstant {
				fuelUsed += int(fuelDelta)
			} else {
				for f := 0; f <= int(fuelDelta); f++ {
					fuelUsed += f
				}
			}
		}

		if leastFuelUsed == 0 {
			leastFuelUsed = fuelUsed
		}

		if fuelUsed < leastFuelUsed {
			leastFuelUsed = fuelUsed
		}
	}

	return leastFuelUsed
}
