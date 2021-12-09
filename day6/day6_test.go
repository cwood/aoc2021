package day6

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/cwood/aoc2021/internal/parse"
	"github.com/stretchr/testify/assert"
)

func TestDay6(t *testing.T) {
	t.Run("Day 6 Part 1 Test Input should be 5934", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		initial := parse.MustAsInt(parse.AsInt(input[0]))

		assert.Equal(t, 26, LanternFishAfterNDays(initial, 18))
		assert.Equal(t, 5934, LanternFishAfterNDays(initial, 80))
	})

	t.Run("Day 6 Part 1", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))
		initial := parse.MustAsInt(parse.AsInt(input[0]))

		t.Logf("Day 6 Part 1: %d", LanternFishAfterNDays(initial, 80))
	})

	t.Run("Day 6 Part 2 Test Input should be 26984457539", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))
		initial := parse.MustAsInt(parse.AsInt(input[0]))

		assert.Equal(t, 26984457539, LanternFishAfterNDays(initial, 256))
	})

	t.Run("Day 6 Part 2", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))
		initial := parse.MustAsInt(parse.AsInt(input[0]))

		t.Logf("Day 6 Part 2: %d", LanternFishAfterNDays(initial, 256))
	})
}
