package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseAsInt(t *testing.T, text []byte) []int {
	ints := make([]int, 0)
	for _, line := range strings.Split(string(text), "\n") {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		require.NoError(t, err)
		ints = append(ints, num)
	}

	return ints
}

func TestDay1(t *testing.T) {
	t.Run("Day1 Test Input should be 7", func(t *testing.T) {
		testInput, err := ioutil.ReadFile("testinput")
		require.NoError(t, err)
		assert.Equal(t, 7, CalculateIncreases(parseAsInt(t, testInput), 1))
	})

	t.Run("Day1 Input 1", func(t *testing.T) {
		input, err := ioutil.ReadFile("input")
		require.NoError(t, err)

		t.Logf("Day 1 Part 1: %d", CalculateIncreases(parseAsInt(t, input), 1))
	})

	t.Run("Day1 Test Input Part 2 should be 5", func(t *testing.T) {
		testInput, err := ioutil.ReadFile("testinput")
		require.NoError(t, err)
		assert.Equal(t, 5, CalculateIncreases(parseAsInt(t, testInput), 3))
	})

	t.Run("Day1 Input Part 2", func(t *testing.T) {
		input, err := ioutil.ReadFile("input")
		require.NoError(t, err)

		t.Logf("Day 1 Part 2: %d", CalculateIncreases(parseAsInt(t, input), 3))
	})
}
