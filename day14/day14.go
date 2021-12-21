package day14

import (
	"strings"
)

func insertI(arr []string, index int, v string) []string {
	if len(arr) == index {
		return append(arr, v)
	}
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = v
	return arr

}

func polymoreStepGen(initial string, rules map[string]string, n int) []string {
	str := strings.Split(initial, "")

step:
	for i := 0; i <= n-1; i++ {

		nstr := make([]string, 0)

		var currC int = 0
		for {
			if currC+1 == len(str) {
				nstr = append(nstr, str[currC])
				str = nstr
				continue step
			}

			r := str[currC] + str[currC+1]
			if newC, ok := rules[r]; ok {
				nstr = append(nstr, str[currC], newC)
				if currC+1 == len(str) {
					nstr = append(nstr, str[currC+1])
				}
			}
			currC++
		}
	}

	return str
}

func PolymoreMostLeastCommon(initial string, rules map[string]string, n int) int {
	tstr := polymoreStepGen(initial, rules, n)

	var counts = make(map[string]int, 0)

	for _, s := range tstr {
		_, ok := counts[s]
		if !ok {
			counts[s] = 0
		}
		counts[s]++
	}

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
