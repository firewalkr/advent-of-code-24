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
00992111777.44.333....5555.6666.....8888..
`)

var previousFileBlockInput = stringToDisk(strings.TrimSpace(`
00992111777.44.333....5555.6666.....8888..
`))

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

func Test_previousFileBlock(t *testing.T) {
	// Test cases
	cases := []struct {
		name           string
		input          []int
		pos            int
		fileID         int
		outputNewPos   int
		outputFileSize int
	}{
		{
			name:           "aoc previous file block sample",
			input:          previousFileBlockInput,
			pos:            41,
			fileID:         8,
			outputNewPos:   36,
			outputFileSize: 4,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got, gotFS := previousFileBlock(c.input, c.pos, c.fileID)
		if got != c.outputNewPos {
			t.Errorf("previousFileBlock(%q) == %v, want %v", c.name, got, c.outputNewPos)
		}
		if gotFS != c.outputFileSize {
			t.Errorf("previousFileBlock(%q) == %v, want %v", c.name, gotFS, c.outputFileSize)
		}
	}
}

func Test_nextEmptyBlock(t *testing.T) {
	// Test cases
	cases := []struct {
		name    string
		disk    []int
		pos     int
		minSize int
		output  int
	}{
		{
			name:    "aoc next empty block sample",
			disk:    moveBlocksInput,
			pos:     0,
			minSize: 2,
			output:  2,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := nextEmptyBlock(c.disk, c.pos, c.minSize)
		if got != c.output {
			t.Errorf("nextEmptyBlock(%q) == %v, want %v", c.name, got, c.output)
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
			output: "00992111777.44.333....5555.6666.....8888..",
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
			output: 2858,
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
