package day10

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	for _, test := range []struct {
		input     string
		corrupt   bool
		badValues string
	}{
		{
			"([])",
			false,
			noCorrupts,
		},
		{
			"(]",
			true,
			"[",
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			isCorrupt, corrupts := isLineCorrupt(test.input)
			assert.Equal(t, test.corrupt, isCorrupt)
			assert.Equal(t, test.badValues, corrupts)
		})
	}

	t.Run("Part 1 Test should be 26397", func(t *testing.T) {
		lines := file.MustLoad(file.LoadAsStrings("testinput"))
		assert.Equal(t, 26397, SyntaxScoringCorrupt(lines))
	})

	t.Run("Part 1", func(t *testing.T) {
		lines := file.MustLoad(file.LoadAsStrings("input"))
		t.Logf("Part 1: %d", SyntaxScoringCorrupt(lines))
	})
}
