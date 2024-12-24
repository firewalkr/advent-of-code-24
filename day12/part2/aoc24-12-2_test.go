package main

import (
	"strings"
	"testing"
)

var aocSample1 = strings.TrimSpace(`
AAAA
BBCD
BBCC
EEEC
`)

var aocSample2 = strings.TrimSpace(`
OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
`)

var aocSample3 = strings.TrimSpace(`
EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
`)

var aocSample4 = strings.TrimSpace(`
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
`)

func Test_scoreRegion(t *testing.T) {
	type args struct {
		grid *Grid
		x, y int
	}
	cases := []struct {
		name   string
		args   args
		output int64
	}{
		{
			name: "aoc sample 0 0",
			args: args{
				grid: stringToGrid(aocSample1),
				x:    0,
				y:    0,
			},
			output: 16,
		},
		{
			name: "aoc sample 1 2",
			args: args{
				grid: stringToGrid(aocSample1),
				x:    1,
				y:    2,
			},
			output: 16,
		},
		{
			name: "aoc sample 3 2",
			args: args{
				grid: stringToGrid(aocSample1),
				x:    3,
				y:    2,
			},
			output: 32,
		},
		{
			name: "aoc sample 3 1",
			args: args{
				grid: stringToGrid(aocSample1),
				x:    3,
				y:    1,
			},
			output: 4,
		},
		{
			name: "aoc sample 1 3",
			args: args{
				grid: stringToGrid(aocSample1),
				x:    1,
				y:    3,
			},
			output: 12,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := scoreRegion(c.args.grid, c.args.x, c.args.y)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}

func Test_scoreMap(t *testing.T) {
	cases := []struct {
		name   string
		grid   *Grid
		output int64
	}{
		{
			name:   "aoc sample 1",
			grid:   stringToGrid(aocSample1),
			output: 80,
		},
		{
			name:   "aoc sample 2",
			grid:   stringToGrid(aocSample2),
			output: 436,
		},
		{
			name:   "aoc sample 3",
			grid:   stringToGrid(aocSample3),
			output: 236,
		},
		{
			name:   "aoc sample 4",
			grid:   stringToGrid(aocSample4),
			output: 368,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := scoreMap(c.grid)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}
