package parse

import (
	"strconv"
	"strings"
)

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
