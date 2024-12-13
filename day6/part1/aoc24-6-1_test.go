package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`)

var willGoUp = strings.TrimSpace(`
..........
....^.....
..........
`)

var willGoRight = strings.TrimSpace(`
..........
....>.....
..........
`)

var willGoDown = strings.TrimSpace(`
..........
....v.....
..........
`)

var willGoLeft = strings.TrimSpace(`
..........
....<.....
..........
`)

var willGoOutOfBounds = strings.TrimSpace(`
..........
.........>
..........
`)

var obstacleUp = strings.TrimSpace(`
..........
....#.....
....^.....
..........
`)

var obstacleRight = strings.TrimSpace(`
..........
....>#....
..........
`)

var obstacleDown = strings.TrimSpace(`
..........
....v.....
....#.....
..........
`)

var obstacleLeft = strings.TrimSpace(`
..........
....#<....
..........
`)

var stuckInLoop = strings.TrimSpace(`
....#.....
...#^#....
....#.....
`)

func Test_run(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "aoc sample",
			input:  testInput,
			output: 41,
		},
	}

	// Execute test cases
	for _, c := range cases {
		grid := NewGrid(c.input)

		err := run(grid)
		if err != nil {
			t.Fatalf("run(%q) error == %v, want nil", c.name, err)
		}

		if len(grid.pastPositions) != c.output {
			t.Errorf("run(%q) == %d, want %d", c.name, len(grid.pastPositions), c.output)
		}
	}
}

func Test_getNextStatus(t *testing.T) {
	// Test cases
	cases := []struct {
		name      string
		input     string
		output    *Status
		outputErr error
	}{
		{
			name:      "will go up",
			input:     willGoUp,
			output:    &Status{Pos{X: 4, Y: 0}, up},
			outputErr: nil,
		},
		{
			name:      "will go right",
			input:     willGoRight,
			output:    &Status{Pos{X: 5, Y: 1}, right},
			outputErr: nil,
		},
		{
			name:      "will go down",
			input:     willGoDown,
			output:    &Status{Pos{X: 4, Y: 2}, down},
			outputErr: nil,
		},
		{
			name:      "will go left",
			input:     willGoLeft,
			output:    &Status{Pos{X: 3, Y: 1}, left},
			outputErr: nil,
		},
		{
			name:      "will go out of bounds",
			input:     willGoOutOfBounds,
			output:    nil,
			outputErr: errOutOfBounds,
		},
		{
			name:      "obstacle up, turn 90 deg clockwise",
			input:     obstacleUp,
			output:    &Status{Pos{X: 5, Y: 2}, right},
			outputErr: nil,
		},
		{
			name:      "obstacle right, turn 90 deg clockwise",
			input:     obstacleRight,
			output:    &Status{Pos{X: 4, Y: 2}, down},
			outputErr: nil,
		},
		{
			name:      "obstacle down, turn 90 deg clockwise",
			input:     obstacleDown,
			output:    &Status{Pos{X: 3, Y: 1}, left},
			outputErr: nil,
		},
		{
			name:      "obstacle left, turn 90 deg clockwise",
			input:     obstacleLeft,
			output:    &Status{Pos{X: 5, Y: 0}, up},
			outputErr: nil,
		},
		{
			name:      "stuck in loop",
			input:     stuckInLoop,
			output:    nil,
			outputErr: errStuckInLoop,
		},
	}

	// Execute test cases
	for _, c := range cases {
		grid := NewGrid(c.input)
		got, err := grid.move()
		if err != c.outputErr {
			t.Errorf("getNextStatus(%q) error == %v, want %v", c.name, err, c.outputErr)
		} else if c.output != nil {
			if got.Pos != c.output.Pos || got.Dir != c.output.Dir {
				t.Errorf("getNextStatus(%q) == %v, want %v", c.name, got, c.output)
			}
		}
	}
}
