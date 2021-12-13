package day11

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay11(t *testing.T) {
	t.Run("Part 1 with 10 steps should be 204", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		octopi := parse.MustMap(parse.AsMap(input))

		assert.Equal(t, 204, OctopusFlashesPerTurn(octopi, 10))
	})

	t.Run("Part 1 with 100 steps should be 1656", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		octopi := parse.MustMap(parse.AsMap(input))

		assert.Equal(t, 1656, OctopusFlashesPerTurn(octopi, 100))
	})
}
