package day14

import (
	"strings"
)

func polymoreStepGen(initial string, rules map[string]string, n int) map[string]int {
	str := strings.Split(initial, "")

	var counts = make(map[string]int)

	for j := 0; j < len(str); j += 1 {
		var pair string
		if j+1 == len(str) {
			continue
		} else {
			pair = str[j] + str[j+1]
		}
		_, ok := counts[pair]
		if !ok {
			counts[pair] = 0
		}
		counts[pair]++
	}

	for i := 0; i <= n-1; i++ {

		ncounts := make(map[string]int)

		for r, v := range counts {
			newC, ok := rules[r]
			if !ok {
				continue
			}
			str := strings.Split(r, "")
			key := str[0] + newC
			okey := newC + str[1]

			_, ok = ncounts[key]
			if !ok {
				ncounts[key] = 0
			}
			_, ok = ncounts[okey]
			if !ok {
				ncounts[okey] = 0
			}
			ncounts[key] += v
			ncounts[okey] += v
		}
		counts = ncounts
	}

	return counts
}

func PolymoreMostLeastCommon(initial string, rules map[string]string, n int) int {
	polyGen := polymoreStepGen(initial, rules, n)

	var counts = make(map[string]int)

	for k, v := range polyGen {
		fk := string(k[0])
		counts[fk] += v
	}

	str := strings.Split(initial, "")

	counts[str[len(str)-1]]++

	var mostCommon, leastCommon = 0, 0

	for _, v := range counts {
		if leastCommon == 0 {
			leastCommon = v
		}
		if v > mostCommon {
			mostCommon = v
		}
		if v < leastCommon {
			leastCommon = v
		}
	}

	return mostCommon - leastCommon
}
