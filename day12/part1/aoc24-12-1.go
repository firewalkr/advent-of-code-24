package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

type Grid struct {
	plots [][]byte
	xSize int
	ySize int
	seen  map[Pos]struct{}
}

func (g *Grid) GetValue(x, y int) byte {
	if x < 0 || y < 0 || x >= g.xSize || y >= g.ySize {
		return 0 // real values are in [A-Z]
	}
	return g.plots[y][x]
}

func (g *Grid) IsVisited(x, y int) bool {
	_, seen := g.seen[Pos{x, y}]
	return seen
}

func (g *Grid) Visit(x, y int) {
	g.seen[Pos{x, y}] = struct{}{}
}

func scoreRegion(g *Grid, x, y int) int64 {
	if g.IsVisited(x, y) {
		return 0
	}

	regionMap := map[Pos]struct{}{}
	regionPerimeter := determineRegion(g, x, y, regionMap)

	return int64(len(regionMap)) * int64(regionPerimeter)
}

func determineRegion(g *Grid, x, y int, regionMap map[Pos]struct{}) (perimeter int) {
	if g.IsVisited(x, y) {
		return
	}

	thisValue := g.GetValue(x, y)
	regionMap[Pos{x, y}] = struct{}{}
	g.Visit(x, y)

	if g.GetValue(x, y-1) == thisValue {
		if !g.IsVisited(x, y-1) {
			perimeter += determineRegion(g, x, y-1, regionMap)
		}
	} else {
		perimeter++
	}

	if g.GetValue(x, y+1) == thisValue {
		if !g.IsVisited(x, y+1) {
			perimeter += determineRegion(g, x, y+1, regionMap)
		}
	} else {
		perimeter++
	}

	if g.GetValue(x-1, y) == thisValue {
		if !g.IsVisited(x-1, y) {
			perimeter += determineRegion(g, x-1, y, regionMap)
		}
	} else {
		perimeter++
	}

	if g.GetValue(x+1, y) == thisValue {
		if !g.IsVisited(x+1, y) {
			perimeter += determineRegion(g, x+1, y, regionMap)
		}
	} else {
		perimeter++
	}

	return
}

func main() {
	input, err := readFile("../aoc24-12-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := stringToGrid(input)

	totalScore := int64(0)
	for y, row := range grid.plots {
		for x := range row {
			totalScore += scoreRegion(grid, x, y)
		}
	}

	t := time.Now()

	fmt.Println(time.Since(t))
	fmt.Println(totalScore)
}

func stringToGrid(input string) *Grid {
	plots := [][]byte{}

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		plots = append(plots, []byte(line))
	}

	xSize := 0
	ySize := len(plots)
	if ySize > 0 {
		xSize = len(plots[0])
	}

	return &Grid{
		plots: plots,
		seen:  map[Pos]struct{}{},
		xSize: xSize,
		ySize: ySize,
	}
}

func readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(fileContent), nil
}
