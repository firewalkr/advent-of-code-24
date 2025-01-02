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

func Test_calcMinPathScore(t *testing.T) {
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
			output: 7036,
		},
		{
			name:   "aoc sample 2",
			args:   args{readGrid(aocSample2)},
			output: 11048,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := calcMinPathScore(c.args.grid)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}
