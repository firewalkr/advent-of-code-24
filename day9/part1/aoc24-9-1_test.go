package main

import (
	"strings"
	"testing"
)

// won't work very well for blocks with ids > 9
// it's just for the test cases
func blocksToStringHelper(blocks []int) string {
	out := strings.Builder{}
	for _, b := range blocks {
		if b == -1 {
			out.WriteString(".")
		} else {
			out.WriteRune(rune(b + 48))
		}

	}

	return out.String()
}

func stringToDisk(input string) []int {
	out := []int{}
	for _, c := range input {
		if c == '.' {
			out = append(out, -1)
		} else {
			out = append(out, int(c)-48)
		}
	}

	return out
}

var mapBlocksSampleInput = strings.TrimSpace(`
2333133121414131402
`)

var moveBlocksInput = stringToDisk(strings.TrimSpace(`
00...111...2...333.44.5555.6666.777.888899
`))

var checksumInput = strings.TrimSpace(`
0099811188827773336446555566..............
`)

func Test_mapBlocks(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "aoc map blocks sample",
			input:  mapBlocksSampleInput,
			output: "00...111...2...333.44.5555.6666.777.888899",
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := mapBlocks(c.input)
		if blocksToStringHelper(got) != c.output {
			t.Errorf("mapBlocks(%q) == %v, want %v", c.name, got, c.output)
		}
	}
}

func Test_moveBlocks(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  []int
		output string
	}{
		{
			name:   "aoc move blocks sample",
			input:  moveBlocksInput,
			output: "0099811188827773336446555566..............",
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := moveBlocks(c.input)
		if blocksToStringHelper(got) != c.output {
			t.Errorf("moveBlocks(%q) == %v, want %v", c.name, got, c.output)
		}
	}
}

func Test_checksum(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "aoc checksum sample",
			input:  checksumInput,
			output: 1928,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := checksum(stringToDisk(c.input))
		if got != c.output {
			t.Errorf("CountAntinodes(%q) == %v, want %v", c.name, got, c.output)
		}
	}
}

func Test_doAll(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "aoc sample",
			input:  mapBlocksSampleInput,
			output: 1928,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := doAll(c.input)
		if got != c.output {
			t.Errorf("doAll(%q) == %v, want %v", c.name, got, c.output)
		}
	}
}
