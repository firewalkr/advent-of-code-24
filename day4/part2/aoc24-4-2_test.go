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

func Test_countCrossedMAS(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "counts all crossed-MAS/SAM",
			input:  testInput,
			output: 9,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := countCrossedMAS(inputAsArrayOfByteArrays(c.input))
		if got != c.output {
			t.Errorf("ReportSafety(%q) == %d, want %d", c.name, got, c.output)
		}
	}
}
