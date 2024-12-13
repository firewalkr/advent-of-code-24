package main

import (
	"strings"
	"testing"
)

var sampleInput = strings.TrimSpace(`
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`)

func Test_countAntinodes(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "aoc sample",
			input:  sampleInput,
			output: 14,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := CountAntinodes(c.input)
		if got != c.output {
			t.Errorf("CountAntinodes(%q) == %v, want %v", c.name, got, c.output)
		}
	}
}
