package day8

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay8(t *testing.T) {
	t.Run("Part 1 Test should be 26", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		inout := parse.AsInOut(input)

		assert.Equal(t, 26, CountUniqueOutput(inout))
	})

	t.Run("Part 1", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))
		inout := parse.AsInOut(input)

		t.Logf("Part 1: %d", CountUniqueOutput(inout))
	})

	t.Run("Part 2 Test should be 61229", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		inout := parse.AsInOut(input)

		assert.Equal(t, 61229, CountUniqueOutput(inout))
	})
}
