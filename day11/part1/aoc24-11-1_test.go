package main

import (
	"strings"
	"testing"
)

var aocSample1 = stringToStones(strings.TrimSpace(`
0 1 10 99 999
`))

var longSample1 = stringToStones(strings.TrimSpace(`
125 17
`))

var longSample2 = stringToStones(strings.TrimSpace(`
253000 1 7
`))

var longSample3 = stringToStones(strings.TrimSpace(`
253 0 2024 14168
`))

var longSample4 = stringToStones(strings.TrimSpace(`
512072 1 20 24 28676032
`))

var longSample5 = stringToStones(strings.TrimSpace(`
512 72 2024 2 0 2 4 2867 6032
`))

var longSample6 = stringToStones(strings.TrimSpace(`
1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32
`))

var longSample7 = stringToStones(strings.TrimSpace(`
2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2
`))

func Test_blink(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		stones []int
		output string
	}{
		{
			name:   "aoc sample 1",
			stones: aocSample1,
			output: "1 2024 1 0 9 9 2021976",
		},
		{
			name:   "long sample 1",
			stones: longSample1,
			output: stonesToString(longSample2),
		},
		{
			name:   "long sample 2",
			stones: longSample2,
			output: stonesToString(longSample3),
		},
		{
			name:   "long sample 3",
			stones: longSample3,
			output: stonesToString(longSample4),
		},
		{
			name:   "long sample 4",
			stones: longSample4,
			output: stonesToString(longSample5),
		},
		{
			name:   "long sample 5",
			stones: longSample5,
			output: stonesToString(longSample6),
		},
		{
			name:   "long sample 6",
			stones: longSample6,
			output: stonesToString(longSample7),
		},
	}

	// Execute test cases
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := blink(c.stones)
			if stonesToString(out) != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}
