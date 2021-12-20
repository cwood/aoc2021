package parse

import (
	"strconv"
	"strings"
)

type InOut struct {
	In  []string
	Out []string
}

func trimStringSlice(m []string) []string {
	var trimmed = make([]string, 0)
	for _, t := range m {
		trimmed = append(trimmed, strings.TrimSpace(t))
	}
	return trimmed
}

func AsInOut(m []string) []InOut {
	inout := make([]InOut, 0)
	for _, l := range m {
		split := strings.SplitN(l, "|", 2)
		i := InOut{
			In:  trimStringSlice(strings.Split(split[0], " ")),
			Out: trimStringSlice(strings.Split(split[1], " ")),
		}
		inout = append(inout, i)
	}
	return inout
}

func AsPointerMap(m []string) map[string]string {
	var pmap = make(map[string]string)

	for _, l := range m {
		prts := strings.SplitN(l, "->", 2)
		pmap[strings.TrimSpace(prts[0])] = strings.TrimSpace(prts[1])
	}

	return pmap
}

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

func MustAsInt(i []int, e error) []int {
	if e != nil {
		panic(e)
	}
	return i
}
