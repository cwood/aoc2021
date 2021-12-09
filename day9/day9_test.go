package day9

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay9(t *testing.T) {
	t.Run("Part 1 Test should be 15", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		points := parse.MustMap(parse.AsMap(input))
		assert.Equal(t, 15, LowpointInMap(points))
	})

	t.Run("Part 1", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))
		points := parse.MustMap(parse.AsMap(input))
		t.Logf("Part 1: %d", LowpointInMap(points))
	})
}
