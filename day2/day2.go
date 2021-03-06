package day2

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type Position string

const (
	Up      Position = "up"
	Forward Position = "forward"
	Down    Position = "down"
)

type Direction struct {
	Position Position
	Number   int
}

func parseAsDirs(t *testing.T, input []byte, out chan Direction) {
	defer close(out)

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		info := strings.SplitN(line, " ", 2)

		n, err := strconv.Atoi(info[1])
		require.NoError(t, err)

		out <- Direction{
			Position: Position(info[0]),
			Number:   n,
		}

	}
}

func Aim(dirs chan Direction) int {
	var (
		forward int
		aim     int
		depth   int
	)

	for dir := range dirs {
		switch dir.Position {
		case Up:
			aim -= dir.Number
		case Down:
			aim += dir.Number
		case Forward:
			forward += dir.Number
			depth += aim * dir.Number
		default:
			panic(fmt.Sprintf("Unknown position %s", dir.Position))
		}
	}

	return depth * forward
}

func Depth(dirs chan Direction) int {
	var (
		depth   int
		forward int
	)

	for dir := range dirs {
		switch dir.Position {
		case Up:
			depth -= dir.Number
		case Down:
			depth += dir.Number
		case Forward:
			forward += dir.Number
		default:
			panic(fmt.Sprintf("Unknown position %s", dir.Position))
		}
	}

	return depth * forward
}
