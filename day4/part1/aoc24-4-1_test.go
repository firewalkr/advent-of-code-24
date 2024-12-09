package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`)

var testInputSingleDiagonalDownRight = strings.TrimSpace(`
......
.X....
..M...
...A..
....S.
......
`)

var testInputSingleDiagonalDownLeft = strings.TrimSpace(`
......
....X.
...M..
..A...
.S....
......
`)

var testInputSingleDiagonalUpRight = strings.TrimSpace(`
......
....S.
...A..
..M...
.X....
......
`)

var testInputSingleDiagonalUpLeft = strings.TrimSpace(`
......
.S....
..A...
...M..
....X.
......
`)

var testInputTwoDiagonalsEndingAtSamePoint = strings.TrimSpace(`
...S...
..A.A..
.M...M.
X.....X
`)

var testInputTwoDiagonalsStartingAtSamePoint = strings.TrimSpace(`
...X...
..M.M..
.A...A.
S.....S
`)

var testInputThreeHorizontalVerticalAndDiagonalStartingAtSamePoint = strings.TrimSpace(`
XMAS
MM..
A.A.
S..S
`)

var testInputAllHorizontalVerticalAndDiagonalStartingAtSamePoint = strings.TrimSpace(`
S..S..S
.A.A.A.
..MMM..
SAMXMAS
..MMM..
.A.A.A.
S..S..S
`)

var testInputAllHorizontalVerticalAndDiagonalEndingAtSamePoint = strings.TrimSpace(`
X..X..X
.M.M.M.
..AAA..
XMASAMX
..AAA..
.M.M.M.
X..X..X
`)

func Test_countXMAS(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "counts all XMAS occurrences, including vertical, diagonal, reversed and overlapping ones",
			input:  testInput,
			output: 18,
		},
		{
			name:   "counts a single diagonal down-right XMAS",
			input:  testInputSingleDiagonalDownRight,
			output: 1,
		},
		{
			name:   "counts a single diagonal down-left XMAS",
			input:  testInputSingleDiagonalDownLeft,
			output: 1,
		},
		{
			name:   "counts a single diagonal up-right XMAS",
			input:  testInputSingleDiagonalUpRight,
			output: 1,
		},
		{
			name:   "counts a single diagonal up-left XMAS",
			input:  testInputSingleDiagonalUpLeft,
			output: 1,
		},
		{
			name:   "counts two diagonals ending at the same point",
			input:  testInputTwoDiagonalsEndingAtSamePoint,
			output: 2,
		},
		{
			name:   "counts two diagonals starting at the same point",
			input:  testInputTwoDiagonalsStartingAtSamePoint,
			output: 2,
		},
		{
			name:   "counts horizontal, vertical and diagonal starting at the same point",
			input:  testInputThreeHorizontalVerticalAndDiagonalStartingAtSamePoint,
			output: 3,
		},
		{
			name:   "counts all horizontal, vertical and diagonal starting at the same point",
			input:  testInputAllHorizontalVerticalAndDiagonalStartingAtSamePoint,
			output: 8,
		},
		{
			name:   "counts all horizontal, vertical and diagonal ending at the same point",
			input:  testInputAllHorizontalVerticalAndDiagonalEndingAtSamePoint,
			output: 8,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := countXMAS(inputAsArrayOfByteArrays(c.input))
		if got != c.output {
			t.Errorf("ReportSafety(%q) == %d, want %d", c.name, got, c.output)
		}
	}
}
