package main

import (
	"strings"
	"testing"
)

var aocSample1 = strings.Split(strings.TrimSpace(`
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`), "\n")

var aocSample2 = strings.Split(strings.TrimSpace(`
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################
`), "\n")

func Test_calcTilesInBestPaths(t *testing.T) {
	type args struct {
		grid *Grid
	}
	cases := []struct {
		name   string
		args   args
		output int
	}{
		{
			name:   "aoc sample 1",
			args:   args{readGrid(aocSample1)},
			output: 45,
		},
		{
			name:   "aoc sample 2",
			args:   args{readGrid(aocSample2)},
			output: 64,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			path, _ := aStar(c.args.grid, func(t Traversal) int {
				return abs(t.pos.x-c.args.grid.endX) + abs(t.pos.y-c.args.grid.endY)
			})
			if len(path) != c.output {
				t.Errorf("Expected %v, but got %v", c.output, len(path))

				posMap := make(map[Pos]struct{})
				for _, pos := range path {
					posMap[pos] = struct{}{}
				}

				t.Log(c.args.grid.PrintWithTilesInBestPaths(posMap))
			}
		})
	}
}
