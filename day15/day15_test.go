package day15

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay15(t *testing.T) {
	t.Run("Part 1 Test should be 40", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		grid := parse.MustMap(parse.AsMap(input))

		assert.Equal(t, 40, LowestTotalRisk(grid))
	})
}
