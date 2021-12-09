package day7

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	t.Run("Part 1 Test should be 37", func(t *testing.T) {
		input := parse.MustAsInt(parse.AsInt(
			file.MustLoad(file.LoadAsStrings("testinput"))[0],
		))

		assert.Equal(t, 37, LeastFuelUsed(input, false))
	})

	t.Run("Part 1", func(t *testing.T) {
		input := parse.MustAsInt(parse.AsInt(
			file.MustLoad(file.LoadAsStrings("input"))[0],
		))

		t.Logf("Part 1: %d", LeastFuelUsed(input, false))
	})

	t.Run("Part 2 Test should be 168", func(t *testing.T) {
		input := parse.MustAsInt(parse.AsInt(
			file.MustLoad(file.LoadAsStrings("testinput"))[0],
		))

		assert.Equal(t, 168, LeastFuelUsed(input, true))
	})

	t.Run("Part 2", func(t *testing.T) {
		input := parse.MustAsInt(parse.AsInt(
			file.MustLoad(file.LoadAsStrings("input"))[0],
		))

		t.Logf("Part 2: %d", LeastFuelUsed(input, true))
	})
}
