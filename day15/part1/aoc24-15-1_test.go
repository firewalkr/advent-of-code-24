package main

import (
	"fmt"
	"strings"
	"testing"
)

var aocMoves = []byte("<^^>>>vv<v>>v<<")

var aocMoveSamples = []*Grid{
	readGrid(strings.Split(strings.TrimSpace(`
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#..@OO.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##..@..#
#...O..#
#.#.O..#
#...O..#
#...O..#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##..@..#
#...O..#
#.#.O..#
#...O..#
#...O..#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##.@...#
#...O..#
#.#.O..#
#...O..#
#...O..#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##.....#
#..@O..#
#.#.O..#
#...O..#
#...O..#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##.....#
#...@O.#
#.#.O..#
#...O..#
#...O..#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##.....#
#....@O#
#.#.O..#
#...O..#
#...O..#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##.....#
#.....O#
#.#.O@.#
#...O..#
#...O..#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########
`), "\n")),
	readGrid(strings.Split(strings.TrimSpace(`
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########
`), "\n")),
}

var aocGpsSampleLarge = strings.Split(strings.TrimSpace(`
##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########
`), "\n")

var aocGpsSampleSmall = strings.Split(strings.TrimSpace(`
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########
`), "\n")

func Test_move(t *testing.T) {
	for i := 0; i < len(aocMoveSamples)-1; i++ {
		t.Run(fmt.Sprintf("Moving from grid %d to grid %d", i, i+1), func(t *testing.T) {
			out := move(aocMoveSamples[i], aocMoves[i])
			if !out.IsEqualTo(aocMoveSamples[i+1]) {
				t.Errorf("Expected:\n%s\nBut got:\n%s\n", aocMoveSamples[i+1], out)
			}
		})
	}
}

func Test_sumGpsCoords(t *testing.T) {
	cases := []struct {
		name   string
		grid   *Grid
		output int
	}{
		{
			name:   "aoc sample large",
			grid:   readGrid(aocGpsSampleLarge),
			output: 10092,
		},
		{
			name:   "aoc sample small",
			grid:   readGrid(aocGpsSampleSmall),
			output: 2028,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := sumGpsCoords(c.grid)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}
