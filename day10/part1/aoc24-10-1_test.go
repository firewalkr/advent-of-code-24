package main

import (
	"strings"
	"testing"
)

var aocSample = stringToTerrain(strings.TrimSpace(`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`))

func Test_scoreAllTrailheads(t *testing.T) {
	// Test cases
	cases := []struct {
		name    string
		terrain [][]int8
		output  int
	}{
		{
			name:    "aoc sample",
			terrain: aocSample,
			output:  36,
		},
	}

	// Execute test cases
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := scoreAllTrailheads(c.terrain)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}
