package parse

import (
	"strconv"
	"strings"
)

func AsMap(m []string) ([][]int, error) {
	var imap = make([][]int, 0)
	for _, l := range m {
		var row = make([]int, 0)
		nums := strings.Split(l, "")
		for _, n := range nums {
			n, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			row = append(row, n)
		}
		imap = append(imap, row)
	}
	return imap, nil
}

func MustMap(i [][]int, e error) [][]int {
	if e != nil {
		panic(e)
	}
	return i
}

func AsInt(line string) ([]int, error) {
	l := make([]int, 0)

	for _, c := range strings.Split(line, ",") {
		n, err := strconv.Atoi(c)
		if err != nil {
			return nil, err
		}
		l = append(l, n)
	}

	return l, nil
}

func MustParse(i []int, e error) []int {
	if e != nil {
		panic(e)
	}
	return i
}
