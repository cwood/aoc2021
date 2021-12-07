package day3

import (
	"testing"

	"github.com/cwood/aoc2021/internal/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay3(t *testing.T) {
	t.Run("Day 3 Part 1 Test Input should be 198", func(t *testing.T) {
		testInput := file.MustLoad(file.LoadAsStrings("testinput"))

		c, err := BinaryDiagnostic(testInput)
		require.NoError(t, err)
		assert.Equal(t, int64(198), c)
	})

	t.Run("Day 3 Part 1 Input", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))

		c, err := BinaryDiagnostic(input)
		require.NoError(t, err)
		t.Logf("Day3 Part 1 Answer: %d", c)
	})

	t.Run("Day 3 Part 2 Test Input should be 230", func(t *testing.T) {
		testInput := file.MustLoad(file.LoadAsStrings("testinput"))

		c, err := BinaryFilter(testInput)
		require.NoError(t, err)
		assert.Equal(t, int64(230), c)
	})

	t.Run("Day 3 Part 2 Input", func(t *testing.T) {
		input := file.MustLoad(file.LoadAsStrings("input"))

		c, err := BinaryFilter(input)
		require.NoError(t, err)
		t.Logf("Day3 Part 2 Answer: %d", c)
	})
}
