package day5

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/stretchr/testify/assert"
)

func TestDay5(t *testing.T) {
	t.Run("Day 5 Part 1 should be 5", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("testinput"))

		var vl = make([]*Line, 0)

		for _, i := range input {
			vl = append(vl, New(i))
		}

		t.Logf("%v", vl)

		assert.Equal(t, 5, LinesIntersect(vl))
	})
}
