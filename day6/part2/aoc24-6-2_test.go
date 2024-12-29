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

var stuckInLoop = strings.TrimSpace(`
....#.....
...#^#....
....#.....
`)

var stuckInLoop2 = strings.TrimSpace(`
.#........
.^......#.
..........
#.........
.......#..
`)

func Test_run(t *testing.T) {
	// Test cases
	cases := []struct {
		name    string
		input   string
		output  int
		wantErr error
	}{
		{
			name:    "aoc sample",
			input:   testInput,
			output:  41,
			wantErr: nil,
		},
		{
			name:    "stuck in loop",
			input:   stuckInLoop,
			wantErr: errStuckInLoop,
		},
		{
			name:    "stuck in loop 2",
			input:   stuckInLoop2,
			wantErr: errStuckInLoop,
		},
	}

	// Execute test cases
	for _, c := range cases {
		grid := NewGrid(c.input)

		botX, botY := grid.getBotPosition()
		directionNode := grid.getValAt(botX, botY)
		grid.setValAt(botX, botY, empty)

		statuses, err := run(grid, botX, botY, directionNode)
		if err != nil {
			if c.wantErr != nil && err == c.wantErr {
				continue
			} else {
				t.Fatalf("run(%q) error == %v, want nil", c.name, err)
			}
		}

		positions := map[Pos]struct{}{}
		for s := range statuses {
			positions[Pos{s.x, s.y}] = struct{}{}
		}

		if len(positions) != c.output {
			t.Errorf("run(%q) == %d, want %d", c.name, len(positions), c.output)
		}
	}
}

// func Test_getNextStatus(t *testing.T) {
// 	// Test cases
// 	cases := []struct {
// 		name      string
// 		input     string
// 		output    *Status
// 		outputErr error
// 	}{
// 		{
// 			name:      "will go up",
// 			input:     willGoUp,
// 			output:    &Status{Pos{X: 4, Y: 0}, up},
// 			outputErr: nil,
// 		},
// 		{
// 			name:      "will go right",
// 			input:     willGoRight,
// 			output:    &Status{Pos{X: 5, Y: 1}, right},
// 			outputErr: nil,
// 		},
// 		{
// 			name:      "will go down",
// 			input:     willGoDown,
// 			output:    &Status{Pos{X: 4, Y: 2}, down},
// 			outputErr: nil,
// 		},
// 		{
// 			name:      "will go left",
// 			input:     willGoLeft,
// 			output:    &Status{Pos{X: 3, Y: 1}, left},
// 			outputErr: nil,
// 		},
// 		{
// 			name:      "will go out of bounds",
// 			input:     willGoOutOfBounds,
// 			output:    nil,
// 			outputErr: errOutOfBounds,
// 		},
// 		{
// 			name:      "obstacle up, turn 90 deg clockwise",
// 			input:     obstacleUp,
// 			output:    &Status{Pos{X: 5, Y: 2}, right},
// 			outputErr: nil,
// 		},
// 		{
// 			name:      "obstacle right, turn 90 deg clockwise",
// 			input:     obstacleRight,
// 			output:    &Status{Pos{X: 4, Y: 2}, down},
// 			outputErr: nil,
// 		},
// 		{
// 			name:      "obstacle down, turn 90 deg clockwise",
// 			input:     obstacleDown,
// 			output:    &Status{Pos{X: 3, Y: 1}, left},
// 			outputErr: nil,
// 		},
// 		{
// 			name:      "obstacle left, turn 90 deg clockwise",
// 			input:     obstacleLeft,
// 			output:    &Status{Pos{X: 5, Y: 0}, up},
// 			outputErr: nil,
// 		},
// 		{
// 			name:      "stuck in loop",
// 			input:     stuckInLoop,
// 			output:    nil,
// 			outputErr: errStuckInLoop,
// 		},
// 	}

// 	// Execute test cases
// 	for _, c := range cases {
// 		grid := NewGrid(c.input)
// 		got, err := grid.move()
// 		if err != c.outputErr {
// 			t.Errorf("getNextStatus(%q) error == %v, want %v", c.name, err, c.outputErr)
// 		} else if c.output != nil {
// 			if got.Pos != c.output.Pos || got.Dir != c.output.Dir {
// 				t.Errorf("getNextStatus(%q) == %v, want %v", c.name, got, c.output)
// 			}
// 		}
// 	}
// }
