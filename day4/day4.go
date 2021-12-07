package day4

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Map    [][]int
	Marked [][]bool
}

func (b *Board) String() string {
	r := "----Start---\n"
	for i, row := range b.Map {
		for c, num := range row {
			if b.Marked[i][c] {
				r += "*"
			}
			r += fmt.Sprintf("%02d ", num)
		}
		r += "\n"
	}
	r += "---Finish---\n"
	return r
}

func (b *Board) IsWinner() (bool, int) {
	//for r, row := range b.Marked {
	//for c, res := range row {
	//if res == true {

	//}
	//}
	//}
	return false, 0
}

func (b *Board) Mark(n int) {
	for r, row := range b.Map {
		for c, num := range row {
			if num == n {
				b.Marked[r][c] = true
				return
			}
		}
	}
}

func NewBoard(lines []string) (*Board, error) {
	b := &Board{
		Map:    make([][]int, 0),
		Marked: make([][]bool, 0),
	}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		nums := strings.Split(line, " ")
		row := make([]int, 0)
		for _, sN := range nums {
			if "" == strings.TrimSpace(sN) {
				continue
			}

			num, err := strconv.Atoi(sN)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}

		b.Map = append(b.Map, row)
		b.Marked = append(b.Marked, make([]bool, len(row)-1))
	}

	return b, nil
}

func Must(b *Board, e error) *Board {
	if e != nil {
		panic(e)
	}
	return b
}

func Bingo(nums []int, boards []*Board) int {
	return 0
}
