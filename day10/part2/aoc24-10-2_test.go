package main

import (
	"strings"
	"testing"
)

var terrainHorizontalOk = stringToTerrain(strings.TrimSpace(`
0123456789
`))

var terrainReverseHorizontalOk = stringToTerrain(strings.TrimSpace(`
9876543210
`))

var terrainVerticalOk = stringToTerrain(strings.TrimSpace(`
0
1
2
3
4
5
6
7
8
9
`))

var terrainReverseVerticalOk = stringToTerrain(strings.TrimSpace(`
9
8
7
6
5
4
3
2
1
0
`))

var vertAndHoriz = stringToTerrain(strings.TrimSpace(`
0123456789
1.........
2.........
3.........
4.........
5.........
6.........
7.........
8.........
9.........
`))

var fCorner = stringToTerrain(strings.TrimSpace(`
0123456789
1.........
2.........
3.........
4.........
56789.....
6.........
7.........
8.........
9.........
`))

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

func Test_countNumTrails(t *testing.T) {
	// Test cases
	cases := []struct {
		name       string
		terrain    [][]int8
		output     int
		xStart     int
		yStart     int
		expected   int8
		numReached int
	}{
		{
			name:       "horizontal ok",
			terrain:    terrainHorizontalOk,
			output:     1,
			xStart:     0,
			yStart:     0,
			expected:   0,
			numReached: 0,
		},
		{
			name:       "reverse horizontal ok",
			terrain:    terrainReverseHorizontalOk,
			output:     1,
			xStart:     9,
			yStart:     0,
			expected:   0,
			numReached: 0,
		},
		{
			name:       "vertical ok",
			terrain:    terrainVerticalOk,
			output:     1,
			xStart:     0,
			yStart:     0,
			expected:   0,
			numReached: 0,
		},
		{
			name:       "reverse vertical ok",
			terrain:    terrainReverseVerticalOk,
			output:     1,
			xStart:     0,
			yStart:     9,
			expected:   0,
			numReached: 0,
		},
		{
			name:       "vertical and horizontal",
			terrain:    vertAndHoriz,
			output:     2,
			xStart:     0,
			yStart:     0,
			expected:   0,
			numReached: 0,
		},
		{
			name:       "f corner",
			terrain:    fCorner,
			output:     3,
			xStart:     0,
			yStart:     0,
			expected:   0,
			numReached: 0,
		},
	}

	// Execute test cases
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := countNumTrails(c.terrain, c.xStart, c.yStart, c.expected, c.numReached)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}

func Test_countAllTrailheads(t *testing.T) {
	// Test cases
	cases := []struct {
		name    string
		terrain [][]int8
		output  int
	}{
		{
			name:    "aoc sample",
			terrain: aocSample,
			output:  81,
		},
	}

	// Execute test cases
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := countAllTrailheads(c.terrain)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}
