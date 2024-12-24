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
			output: 40,
		},
		{
			name: "aoc sample 1 2",
			args: args{
				grid: stringToGrid(aocSample1),
				x:    1,
				y:    2,
			},
			output: 32,
		},
		{
			name: "aoc sample 3 2",
			args: args{
				grid: stringToGrid(aocSample1),
				x:    3,
				y:    2,
			},
			output: 40,
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
			output: 24,
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
