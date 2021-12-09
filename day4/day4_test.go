package day4

import (
	"fmt"
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	t.Run("Day 4 Part 1 Test Should be 4512", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		bingoNums := parse.MustAsInt(parse.AsInt(input[0]))

		var boards []*Board
		var rawBoards = input[1 : len(input)-1]

		for i := 0; i <= len(rawBoards)-1; i += 5 {
			boards = append(boards, Must(NewBoard(
				rawBoards[i:i+5],
			)))
		}

		assert.Equal(t, 4512, Bingo(bingoNums, boards), fmt.Sprintf("%v", boards))
	})

	t.Run("Day 4 Part 1", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))
		bingoNums := parse.MustAsInt(parse.AsInt(input[0]))

		var boards []*Board
		var rawBoards = input[1 : len(input)-1]

		for i := 0; i <= len(rawBoards)-1; i += 5 {
			boards = append(boards, Must(NewBoard(
				rawBoards[i:i+5],
			)))
		}

		t.Logf("Day 4 Part 1: %d", Bingo(bingoNums, boards))
	})

	t.Run("Day 4 Part 2 Test Should be 1924", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		bingoNums := parse.MustAsInt(parse.AsInt(input[0]))

		var boards []*Board
		var rawBoards = input[1 : len(input)-1]

		for i := 0; i <= len(rawBoards)-1; i += 5 {
			boards = append(boards, Must(NewBoard(
				rawBoards[i:i+5],
			)))
		}

		assert.Equal(t, 1924, BingoLastWinner(bingoNums, boards), fmt.Sprintf("%v", boards))
	})

	t.Run("Day 4 Part 2", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))
		bingoNums := parse.MustAsInt(parse.AsInt(input[0]))

		var boards []*Board
		var rawBoards = input[1 : len(input)-1]

		for i := 0; i <= len(rawBoards)-1; i += 5 {
			boards = append(boards, Must(NewBoard(
				rawBoards[i:i+5],
			)))
		}

		t.Logf("Day 4 Part 2: %d", BingoLastWinner(bingoNums, boards))
	})
}
