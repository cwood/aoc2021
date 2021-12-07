package day3

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay3(t *testing.T) {
	t.Run("Day 3 Part 1 Test Input should be 198", func(t *testing.T) {
		testInput, err := ioutil.ReadFile("testinput")
		require.NoError(t, err)

		c, err := BinaryDiagnostic(strings.Split(string(testInput), "\n"))
		require.NoError(t, err)
		assert.Equal(t, int64(198), c)
	})

	t.Run("Day 3 Part 1 Input", func(t *testing.T) {
		input, err := ioutil.ReadFile("input")
		require.NoError(t, err)

		c, err := BinaryDiagnostic(strings.Split(string(input), "\n"))
		require.NoError(t, err)
		t.Logf("Day3 Part 1 Answer: %d", c)
	})

	t.Run("Day 3 Part 2 Test Input should be 230", func(t *testing.T) {
		testInput, err := ioutil.ReadFile("testinput")
		require.NoError(t, err)

		c, err := BinaryFilter(strings.Split(string(testInput), "\n"))
		require.NoError(t, err)
		assert.Equal(t, int64(230), c)
	})
}
