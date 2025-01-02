package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const Wall = byte('#')
const Start = byte('S')
const Empty = byte('.')
const End = byte('E')

type Grid struct {
	grid           [][]byte
	startX, startY int
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

func (g *Grid) String() string {
	sb := strings.Builder{}
	for _, row := range g.grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}

	return strings.TrimSpace(sb.String())
}

func calcMinPathScore(g *Grid) int {
	return calcMinPathScoreHelper(g, g.startX, g.startY, East, 0, SeenWithScore{})
}

type Pos struct {
	x, y int
}

type SeenWithScore map[Pos]int

func calcMinPathScoreHelper(g *Grid, x, y int, direction Direction, score int, seenWithScore SeenWithScore) int {
	if g.GetValue(x, y) == Wall {
		return -1
	}
	if g.GetValue(x, y) == End {
		return score
	}

	if priorScore, seen := seenWithScore[Pos{x, y}]; seen {
		if priorScore <= score+1 {
			return -1
		}
	}

	seenWithScore[Pos{x, y}] = score + 1

	minDownrange := math.MaxInt

	deltaX, deltaY := direction.deltas()

	downrangeScore := calcMinPathScoreHelper(g, x+deltaX, y+deltaY, direction, score+1, seenWithScore)
	if downrangeScore != -1 {
		minDownrange = min(minDownrange, downrangeScore)
	}

	direction90CW := direction.rotate90CW()
	deltaX, deltaY = direction90CW.deltas()

	downrangeScore = calcMinPathScoreHelper(g, x+deltaX, y+deltaY, direction90CW, score+1001, seenWithScore)
	if downrangeScore != -1 {
		minDownrange = min(minDownrange, downrangeScore)
	}

	direction90CCW := direction.rotate90CCW()
	deltaX, deltaY = direction90CCW.deltas()

	downrangeScore = calcMinPathScoreHelper(g, x+deltaX, y+deltaY, direction90CCW, score+1001, seenWithScore)
	if downrangeScore != -1 {
		minDownrange = min(minDownrange, downrangeScore)
	}

	return minDownrange
}

func main() {
	gridStr, err := readFile("../aoc24-16-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := readGrid(gridStr)

	fmt.Println(calcMinPathScore(grid))
}

func readGrid(input []string) *Grid {
	grid := &Grid{}

	for _, line := range input {
		grid.grid = append(grid.grid, []byte(line))
	}

	startX, startY := grid.FindStartPos()

	grid.startX = startX
	grid.startY = startY

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
