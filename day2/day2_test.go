package day2

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay2(t *testing.T) {
	t.Run("Day 2 Part 1 should be 150", func(t *testing.T) {
		testInput, err := ioutil.ReadFile("testinput")
		require.NoError(t, err)

		dirs := make(chan Direction)
		go parseAsDirs(t, testInput, dirs)

		assert.Equal(t, 150, Depth(dirs))
	})

	t.Run("Day 2 Part 1", func(t *testing.T) {
		testInput, err := ioutil.ReadFile("input")
		require.NoError(t, err)

		dirs := make(chan Direction)
		go parseAsDirs(t, testInput, dirs)

		t.Logf("Day 2 Part 1 Answer: %d", Depth(dirs))
	})

	t.Run("Day 2 Part 2 should be 900", func(t *testing.T) {
		testInput, err := ioutil.ReadFile("testinput")
		require.NoError(t, err)

		dirs := make(chan Direction)
		go parseAsDirs(t, testInput, dirs)

		assert.Equal(t, 900, Aim(dirs))
	})

	t.Run("Day 2 Part 2", func(t *testing.T) {
		testInput, err := ioutil.ReadFile("input")
		require.NoError(t, err)

		dirs := make(chan Direction)
		go parseAsDirs(t, testInput, dirs)

		t.Logf("Day 2 Part 2 Answer: %d", Aim(dirs))
	})
}
