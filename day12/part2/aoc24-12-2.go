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
	seen  visitMap
}

func (g *Grid) GetValue(x, y int) byte {
	if x < 0 || y < 0 || x >= g.xSize || y >= g.ySize {
		return 0 // real values are in [A-Z]
	}
	return g.plots[y][x]
}

func (g *Grid) IsVisited(x, y int) bool {
	return g.seen.IsVisited(x, y)
}

func (g *Grid) IsLeftEdge(x, y int) bool {
	return g.GetValue(x, y) != g.GetValue(x-1, y)
}

func (g *Grid) IsRightEdge(x, y int) bool {
	return g.GetValue(x, y) != g.GetValue(x+1, y)
}

func (g *Grid) IsTopEdge(x, y int) bool {
	return g.GetValue(x, y) != g.GetValue(x, y-1)
}

func (g *Grid) IsBottomEdge(x, y int) bool {
	return g.GetValue(x, y) != g.GetValue(x, y+1)
}

func (g *Grid) Visit(x, y int) {
	g.seen.Visit(x, y)
}

type maps struct {
	regionMap                            visitMap
	leftEdgeVisitedMap                   visitMap
	rightEdgeVisitedMap                  visitMap
	bottomEdgeVisitedMap                 visitMap
	topEdgeVisitedMap                    visitMap
	numLeft, numRight, numTop, numBottom int
}

type visitMap map[Pos]struct{}

func (v visitMap) IsVisited(x, y int) bool {
	_, seen := v[Pos{x, y}]
	return seen
}

func (v visitMap) Visit(x, y int) {
	v[Pos{x, y}] = struct{}{}
}

func scoreRegion(g *Grid, x, y int) int64 {
	if g.IsVisited(x, y) {
		return 0
	}

	maps := maps{
		regionMap:            visitMap{},
		leftEdgeVisitedMap:   visitMap{},
		rightEdgeVisitedMap:  visitMap{},
		bottomEdgeVisitedMap: visitMap{},
		topEdgeVisitedMap:    visitMap{},
	}
	determineRegion(g, x, y, &maps)

	//FIXME!!!!!
	return int64(len(maps.regionMap)) * int64(maps.numBottom+maps.numLeft+maps.numTop+maps.numRight)
}

// The strategy for counting sides is simple but brute-force'ish
// and the code could be DRYer. But hey, it ran correctly the first time!
//
// This keeps 4 maps of plots that were visited for left/right/bottom/top side checks.
//
// It goes like this (e.g. for counting top sides):
//
// - if the current plot has NOT been visited on a top side check
//
// - and the current plot is a top side (isTopEdge)
//
// - walk to the left while plots belong to same region and are top sides
//   - mark them "visited" for top side checks as we go
//
// - walk to the right while plots belong to same region and are top sides
//   - mark them "visited" for top side checks as we go
//
// repeat the above for counting left/bottom/right sides
func determineRegion(g *Grid, x, y int, maps *maps) {
	if g.IsVisited(x, y) {
		return
	}

	thisValue := g.GetValue(x, y)

	if !maps.topEdgeVisitedMap.IsVisited(x, y) && g.IsTopEdge(x, y) {
		maps.numTop++
		xl := x - 1
		for {
			if g.GetValue(xl, y) == thisValue && g.IsTopEdge(xl, y) {
				maps.topEdgeVisitedMap.Visit(xl, y)
				xl--
			} else {
				break
			}
		}
		xr := x + 1
		for {
			if g.GetValue(xr, y) == thisValue && g.IsTopEdge(xr, y) {
				maps.topEdgeVisitedMap.Visit(xr, y)
				xr++
			} else {
				break
			}
		}
	}

	if !maps.bottomEdgeVisitedMap.IsVisited(x, y) && g.IsBottomEdge(x, y) {
		maps.numBottom++
		xl := x - 1
		for {
			if g.GetValue(xl, y) == thisValue && g.IsBottomEdge(xl, y) {
				maps.bottomEdgeVisitedMap.Visit(xl, y)
				xl--
			} else {
				break
			}
		}
		xr := x + 1
		for {
			if g.GetValue(xr, y) == thisValue && g.IsBottomEdge(xr, y) {
				maps.bottomEdgeVisitedMap.Visit(xr, y)
				xr++
			} else {
				break
			}
		}
	}

	if !maps.leftEdgeVisitedMap.IsVisited(x, y) && g.IsLeftEdge(x, y) {
		maps.numLeft++
		yu := y - 1
		for {
			if g.GetValue(x, yu) == thisValue && g.IsLeftEdge(x, yu) {
				maps.leftEdgeVisitedMap.Visit(x, yu)
				yu--
			} else {
				break
			}
		}
		yd := y + 1
		for {
			if g.GetValue(x, yd) == thisValue && g.IsLeftEdge(x, yd) {
				maps.leftEdgeVisitedMap.Visit(x, yd)
				yd++
			} else {
				break
			}
		}
	}

	if !maps.rightEdgeVisitedMap.IsVisited(x, y) && g.IsRightEdge(x, y) {
		maps.numRight++
		yu := y - 1
		for {
			if g.GetValue(x, yu) == thisValue && g.IsRightEdge(x, yu) {
				maps.rightEdgeVisitedMap.Visit(x, yu)
				yu--
			} else {
				break
			}
		}
		yd := y + 1
		for {
			if g.GetValue(x, yd) == thisValue && g.IsRightEdge(x, yd) {
				maps.rightEdgeVisitedMap.Visit(x, yd)
				yd++
			} else {
				break
			}
		}
	}

	maps.regionMap[Pos{x, y}] = struct{}{}
	g.Visit(x, y)

	if g.GetValue(x, y-1) == thisValue && !g.IsVisited(x, y-1) {
		determineRegion(g, x, y-1, maps)
	}

	if g.GetValue(x, y+1) == thisValue && !g.IsVisited(x, y+1) {
		determineRegion(g, x, y+1, maps)
	}

	if g.GetValue(x-1, y) == thisValue && !g.IsVisited(x-1, y) {
		determineRegion(g, x-1, y, maps)
	}

	if g.GetValue(x+1, y) == thisValue && !g.IsVisited(x+1, y) {
		determineRegion(g, x+1, y, maps)
	}
}

func scoreMap(g *Grid) int64 {
	totalScore := int64(0)
	for y, row := range g.plots {
		for x := range row {
			totalScore += scoreRegion(g, x, y)
		}
	}

	return totalScore
}

func main() {
	input, err := readFile("../aoc24-12-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := stringToGrid(input)

	totalScore := scoreMap(grid)

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
