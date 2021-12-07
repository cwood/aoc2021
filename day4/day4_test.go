package day4

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay4(t *testing.T) {
	t.Run("Day 4 Part 1 Test Should be 4512", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		bingoNums := parse.MustParse(parse.AsInt(input[0]))

		var boards []*Board
		var rawBoards = input[1 : len(input)-1]

		for i := 0; i <= len(rawBoards)-1; i += 5 {
			boards = append(boards, Must(NewBoard(
				rawBoards[i:i+5],
			)))
		}

		t.Logf("\n%v", boards)

		assert.Equal(t, 4512, Bingo(bingoNums, boards))
	})
}
