package day14

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay14(t *testing.T) {
	t.Run("Day 14 Part 1 Test should be 1588", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		rules := parse.AsPointerMap(input[1:])

		assert.Equal(t, 1588, PolymoreMostLeastCommon(input[0], rules, 10))
	})

	t.Run("Day 14 Part 1", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))
		rules := parse.AsPointerMap(input[1:])

		t.Logf("Part 1: %d", PolymoreMostLeastCommon(input[0], rules, 10))
	})

	t.Run("Day 14 Part 2 Test should be 2188189693529", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		rules := parse.AsPointerMap(input[1:])

		assert.Equal(t, 2188189693529, PolymoreMostLeastCommon(input[0], rules, 40))
	})

}
