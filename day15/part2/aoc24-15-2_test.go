package main

import (
	"strings"
	"testing"
)

var aocGridSample = strings.Split(strings.TrimSpace(`
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########
`), "\n")

var aocWideGridResultSample = strings.Split(strings.TrimSpace(`
####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##....[]@.....[]..##
##[]##....[]......##
##[]....[]....[]..##
##..[][]..[]..[][]##
##........[]......##
####################
`), "\n")

var moveLeftSample = strings.Split(strings.TrimSpace(`
####################
##................##
##..[][]@.........##
##................##
####################
`), "\n")

var moveLeftSampleResult = strings.Split(strings.TrimSpace(`
####################
##................##
##.[][]@..........##
##................##
####################
`), "\n")

var moveLeftSample2 = strings.Split(strings.TrimSpace(`
####################
##................##
##[][]@...........##
##................##
####################
`), "\n")

var moveRightSample = strings.Split(strings.TrimSpace(`
####################
##................##
##.........@[][]..##
##................##
####################
`), "\n")

var moveRightSampleResult = strings.Split(strings.TrimSpace(`
####################
##................##
##..........@[][].##
##................##
####################
`), "\n")

var moveRightSample2 = strings.Split(strings.TrimSpace(`
####################
##................##
##...........@[][]##
##................##
####################
`), "\n")

var moveUpBasic = strings.Split(strings.TrimSpace(`
####################
##................##
##..........[]....##
##...........@....##
####################
`), "\n")

var moveUpBasicResult = strings.Split(strings.TrimSpace(`
####################
##..........[]....##
##...........@....##
##................##
####################
`), "\n")

var moveUpComplex = strings.Split(strings.TrimSpace(`
####################
##................##
##..........[][]..##
##...........[]...##
##..........[][]..##
##...........@....##
##................##
####################
`), "\n")

var moveUpComplexResult = strings.Split(strings.TrimSpace(`
####################
##..........[][]..##
##...........[]...##
##..........[]....##
##...........@[]..##
##................##
##................##
####################
`), "\n")

func Test_move(t *testing.T) {
	type args struct {
		grid *Grid
		move byte
	}
	cases := []struct {
		name   string
		args   args
		output *Grid
	}{
		{
			name: "move boxes left",
			args: args{
				grid: readGrid(moveLeftSample),
				move: byte('<'),
			},
			output: readGrid(moveLeftSampleResult),
		},
		{
			name: "want to move left but it's a wall, so don't move",
			args: args{
				grid: readGrid(moveLeftSample2),
				move: byte('<'),
			},
			output: readGrid(moveLeftSample2),
		},
		{
			name: "move boxes right",
			args: args{
				grid: readGrid(moveRightSample),
				move: byte('>'),
			},
			output: readGrid(moveRightSampleResult),
		},
		{
			name: "want to move right but it's a wall, so don't move",
			args: args{
				grid: readGrid(moveRightSample2),
				move: byte('>'),
			},
			output: readGrid(moveRightSample2),
		},
		{
			name: "move up",
			args: args{
				grid: readGrid(moveUpBasic),
				move: byte('^'),
			},
			output: readGrid(moveUpBasicResult),
		},
		{
			name: "move up complex",
			args: args{
				grid: readGrid(moveUpComplex),
				move: byte('^'),
			},
			output: readGrid(moveUpComplexResult),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := move(c.args.grid, c.args.move)
			if !out.IsEqualTo(c.output) {
				t.Errorf("Expected\n%s, but got\n%s", c.output, out)
			}
		})
	}
}

// func Test_move(t *testing.T) {
// 	for i := 0; i < len(aocMoveSamples)-1; i++ {
// 		t.Run(fmt.Sprintf("Moving from grid %d to grid %d", i, i+1), func(t *testing.T) {
// 			out := move(aocMoveSamples[i], aocMoves[i])
// 			if !out.IsEqualTo(aocMoveSamples[i+1]) {
// 				t.Errorf("Expected:\n%s\nBut got:\n%s\n", aocMoveSamples[i+1], out)
// 			}
// 		})
// 	}
// }

// func Test_sumGpsCoords(t *testing.T) {
// 	cases := []struct {
// 		name   string
// 		grid   *Grid
// 		output int
// 	}{
// 		{
// 			name:   "aoc sample large",
// 			grid:   readGrid(aocGpsSampleLarge),
// 			output: 10092,
// 		},
// 		{
// 			name:   "aoc sample small",
// 			grid:   readGrid(aocGpsSampleSmall),
// 			output: 2028,
// 		},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			out := sumGpsCoords(c.grid)
// 			if out != c.output {
// 				t.Errorf("Expected %v, but got %v", c.output, out)
// 			}
// 		})
// 	}
// }

func Test_widenGrid(t *testing.T) {
	cases := []struct {
		name   string
		grid   *Grid
		output *Grid
	}{
		{
			name:   "widen sample 1",
			grid:   readGrid(aocGridSample),
			output: readGrid(aocWideGridResultSample),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := widenGrid(c.grid)
			if !out.IsEqualTo(c.output) {
				t.Errorf("Expected\n%s, but got\n%s", c.output, out)
			}
		})
	}
}
