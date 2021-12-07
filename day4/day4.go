package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func allAreTrue(i []bool) bool {
	for _, b := range i {
		if !b {
			return false
		}
	}
	return true
}

func sum(n []int) int {
	var s int
	for _, add := range n {
		s += add
	}
	return s
}

type Board struct {
	Map    [][]int
	Marked [][]bool
}

func (b *Board) String() string {
	r := "----Start---\n"
	for i, row := range b.Map {
		for c, num := range row {
			if b.Marked[i][c] {
				r += "[*]"
			} else {
				r += "[ ]"
			}
			r += fmt.Sprintf("%02d ", num)
		}
		r += "\n"
	}
	r += "---Finish---\n"
	return r
}

func (b *Board) sumUnmarked() int {
	var sum int

	for r, row := range b.Marked {
		for c, res := range row {
			if !res {
				sum += b.Map[r][c]
			}
		}
	}
	return sum
}

func (b *Board) columnAsSlice(c int) []bool {
	var r = make([]bool, 0)
	for i, _ := range b.Marked {
		r = append(r, b.Marked[i][c])
	}
	return r
}

func (b *Board) IsWinner() (bool, int) {
	for _, row := range b.Marked {
		for c, res := range row {
			if res == true {
				col := b.columnAsSlice(c)

				if allAreTrue(row) || allAreTrue(col) {
					return true, b.sumUnmarked()
				}
			}
		}
	}
	return false, 0
}

func (b *Board) Mark(n int) bool {
	for r, row := range b.Map {
		for c, num := range row {
			if num == n {
				b.Marked[r][c] = true
				return true
			}
		}
	}
	return false
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
			if strings.TrimSpace(sN) == "" {
				continue
			}

			num, err := strconv.Atoi(sN)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}

		b.Map = append(b.Map, row)
		b.Marked = append(b.Marked, make([]bool, len(row)))
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
	for _, bingoNum := range nums {
		for i := 0; i <= len(boards)-1; i++ {
			if marked := boards[i].Mark(bingoNum); marked {
				if hasWin, rowSum := boards[i].IsWinner(); hasWin {
					return rowSum * bingoNum
				}
			}
		}
	}
	return 0
}

func BingoLastWinner(nums []int, boards []*Board) int {
	var latestWin int

	var boardWins = make(map[int]bool)

	for _, bingoNum := range nums {
		for i := 0; i <= len(boards)-1; i++ {
			if _, ok := boardWins[i]; ok {
				continue
			}
			if marked := boards[i].Mark(bingoNum); marked {
				if hasWin, rowSum := boards[i].IsWinner(); hasWin {
					boardWins[i] = true
					latestWin = rowSum * bingoNum
				}
			}
		}
	}

	return latestWin
}
