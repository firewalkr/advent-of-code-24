package main

import (
	"fmt"
	"math"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

// Some differences from part 1:
//
// - 90CW and 90CCW rotations are now its own item in the "traversal stack"
// - this is because we need to keep track of the direction we're facing when
// arriving at a given position to properly calculate the score
// - the way the structs are being populated is too heavy and should be refactored
// - it was much faster to limit recursion to the score of the overall best path
// which we know from part 1. It's the equivalent of running part 1, which runs
// instantaneously, followed by the part 2 code.
//

const Wall = byte('#')
const Start = byte('S')
const Empty = byte('.')
const End = byte('E')

type Grid struct {
	grid           [][]byte
	startX, startY int
	endX, endY     int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) deltas() (x int, y int) {
	switch d {
	case North:
		return 0, -1
	case East:
		return 1, 0
	case South:
		return 0, 1
	case West:
		return -1, 0
	}

	panic("invalid direction")
}

func (d Direction) rotate90CW() Direction {
	return (d + 1) % 4
}

func (d Direction) rotate90CCW() Direction {
	return (d + 3) % 4
}

func (g *Grid) GetValue(x, y int) byte {
	return g.grid[y][x]
}

func (g *Grid) SetValue(xPos, yPos int, b byte) {
	g.grid[yPos][xPos] = b
}

func (g *Grid) FindStartPos() (xPos, yPos int) {
	for y := range g.grid {
		for x := range g.grid[y] {
			if g.GetValue(x, y) == Start {
				return x, y
			}
		}
	}

	return -1, -1
}

func (g *Grid) FindEndPos() (xPos, yPos int) {
	for y := range g.grid {
		for x := range g.grid[y] {
			if g.GetValue(x, y) == End {
				return x, y
			}
		}
	}

	return -1, -1
}

func (g *Grid) String() string {
	sb := strings.Builder{}
	for _, row := range g.grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}

	return strings.TrimSpace(sb.String())
}

func (g *Grid) PrintWithTilesInBestPaths(tilesInBestPaths map[Pos]struct{}) string {
	sb := strings.Builder{}
	for y, row := range g.grid {
		for x, b := range row {
			if _, ok := tilesInBestPaths[Pos{x, y}]; ok {
				sb.WriteString("O")
			} else {
				sb.WriteByte(b)
			}
		}
		sb.WriteString("\n")
	}

	return strings.TrimSpace(sb.String())
}

var currentPath Path

func calcTilesInBestPaths(g *Grid) int {
	now := time.Now().UTC()

	rows := len(g.grid)
	cols := len(g.grid[0])

	currentPath = make(Path, rows*cols)

	bestScoreTraversed := BestScoreTraversed{score: math.MaxInt}
	bestScore := calcTilesInBestPathsHelper(g, g.startX, g.startY, East, 0, SeenWithScore{}, 0, &bestScoreTraversed)

	fmt.Println("best score:", bestScore)

	positions := map[Pos]struct{}{}

	for t := range bestScoreTraversed.traversalsWithScore {
		positions[t.pos] = struct{}{}
	}

	fmt.Println(g.PrintWithTilesInBestPaths(positions))
	fmt.Println(time.Now().UTC().Sub(now).String())

	return len(positions)
}

type Pos struct {
	x, y int
}

type Traversal struct {
	pos       Pos
	direction Direction
}

type TraversalWithScore struct {
	Traversal
	score int
}

type SeenWithScore map[Traversal]int

type Path []TraversalWithScore

type Paths []Path

type BestScoreTraversed struct {
	score               int
	traversalsWithScore map[TraversalWithScore]struct{}
}

var numNodeVisits int

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func calcTilesInBestPathsHelper(
	g *Grid,
	x, y int,
	direction Direction,
	score int,
	seenWithScore SeenWithScore,
	numNode int,
	bestScoreTraversed *BestScoreTraversed) int {
	numNodeVisits++

	thisTraversal := TraversalWithScore{Traversal{Pos{x, y}, direction}, score}
	currentPath[numNode] = thisTraversal

	if g.GetValue(x, y) == End {
		if score <= bestScoreTraversed.score {
			if score < bestScoreTraversed.score {
				bestScoreTraversed.traversalsWithScore = make(map[TraversalWithScore]struct{})
			}

			bestScoreTraversed.score = score
			for i := 0; i <= numNode; i++ {
				bestScoreTraversed.traversalsWithScore[currentPath[i]] = struct{}{}
			}
		}

		return score
	}

	// if score >= 135512 {
	// 	return -1
	// }

	if priorScore, seen := seenWithScore[thisTraversal.Traversal]; seen {
		if priorScore < score {
			return -1
		}
		if _, seenInBestPath := bestScoreTraversed.traversalsWithScore[thisTraversal]; seenInBestPath {
			for i := 0; i <= numNode; i++ {
				bestScoreTraversed.traversalsWithScore[currentPath[i]] = struct{}{}
			}
			return bestScoreTraversed.score
		}
	}

	seenWithScore[thisTraversal.Traversal] = score

	minDownrange := math.MaxInt

	deltaX, deltaY := direction.deltas()

	if g.GetValue(x+deltaX, y+deltaY) != Wall && score+1 < bestScoreTraversed.score {
		downrangeScore := calcTilesInBestPathsHelper(g, x+deltaX, y+deltaY, direction, score+1, seenWithScore, numNode+1, bestScoreTraversed)
		if downrangeScore != -1 {
			minDownrange = min(downrangeScore, minDownrange)
		}
	}

	direction90CW := direction.rotate90CW()
	deltaX, deltaY = direction90CW.deltas()

	if g.GetValue(x+deltaX, y+deltaY) != Wall && score+1000 < bestScoreTraversed.score {
		downrangeScore := calcTilesInBestPathsHelper(g, x, y, direction90CW, score+1000, seenWithScore, numNode+1, bestScoreTraversed)
		if downrangeScore != -1 {
			minDownrange = min(downrangeScore, minDownrange)
		}
	}

	direction90CCW := direction.rotate90CCW()
	deltaX, deltaY = direction90CCW.deltas()

	if g.GetValue(x+deltaX, y+deltaY) != Wall && score+1000 < bestScoreTraversed.score {
		downrangeScore := calcTilesInBestPathsHelper(g, x, y, direction90CCW, score+1000, seenWithScore, numNode+1, bestScoreTraversed)
		if downrangeScore != -1 {
			minDownrange = min(downrangeScore, minDownrange)
		}
	}

	return minDownrange
}

func main() {
	// Create CPU profile file
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Start CPU profiling
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	gridStr, err := readFile("../aoc24-16-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := readGrid(gridStr)

	fmt.Println(calcTilesInBestPaths(grid))
	fmt.Println(numNodeVisits)
}

func readGrid(input []string) *Grid {
	grid := &Grid{}

	for _, line := range input {
		grid.grid = append(grid.grid, []byte(line))
	}

	grid.startX, grid.startY = grid.FindStartPos()
	grid.endX, grid.endY = grid.FindEndPos()

	return grid
}

func readFile(filename string) (gridStr []string, err error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	fileLines := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")

	i := 0
	for _, line := range fileLines {
		gridStr = append(gridStr, line)
		i++
	}

	return gridStr, nil
}
