package main

import (
	"fmt"
	"os"
	"strings"
)

const Wall = byte('#')
const Box = byte('O')
const Robot = byte('@')
const Empty = byte('.')
const OutOfBounds = byte(0)

const Up = byte('^')
const Down = byte('v')
const Left = byte('<')
const Right = byte('>')

type Grid struct {
	grid [][]byte
	xLen int
	yLen int
}

func (g *Grid) GetValue(x, y int) byte {
	if x < 0 || y < 0 || x >= g.xLen || y >= g.yLen {
		return OutOfBounds
	}
	return g.grid[y][x]
}

func (g *Grid) SetValue(xPos, yPos int, b byte) {
	g.grid[yPos][xPos] = b
}

func (g *Grid) IsEqualTo(g2 *Grid) bool {
	if g.xLen != g2.xLen || g.yLen != g2.yLen {
		return false
	}

	for y := range g.grid {
		for x := range g.grid[y] {
			if g.GetValue(x, y) != g2.GetValue(x, y) {
				return false
			}
		}
	}

	return true
}

func (g *Grid) FindRobot() (xPos, yPos int) {
	for y := range g.grid {
		for x := range g.grid[y] {
			if g.GetValue(x, y) == Robot {
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

func (g *Grid) Clone() *Grid {
	newGrid := &Grid{
		xLen: g.xLen,
		yLen: g.yLen,
	}

	for _, row := range g.grid {
		newGrid.grid = append(newGrid.grid, append([]byte{}, row...))
	}

	return newGrid
}

func move(g *Grid, move byte) *Grid {
	newGrid := g.Clone()

	botX, botY := newGrid.FindRobot()

	var xDelta = 0
	var yDelta = 0
	if move == Left {
		xDelta = -1
	} else if move == Right {
		xDelta = 1
	} else if move == Up {
		yDelta = -1
	} else if move == Down {
		yDelta = 1
	} else {
		panic("invalid move")
	}

	newGrid.moveThings(botX, botY, xDelta, yDelta, Empty)

	return newGrid
}

func (g *Grid) moveThings(xPos, yPos, xDelta, yDelta int, thing byte) (behindBoxes byte) {
	this := g.GetValue(xPos, yPos)
	next := g.GetValue(xPos+xDelta, yPos+yDelta)
	if next == Wall {
		return Wall
	} else if next == Empty {
		g.SetValue(xPos+xDelta, yPos+yDelta, this)
		g.SetValue(xPos, yPos, thing)
		return Empty
	} // shouldn't need to check for out of bounds since the grid is surrounded by walls

	// else if Box
	behind := g.moveThings(xPos+xDelta, yPos+yDelta, xDelta, yDelta, this)
	if behind == Wall {
		return Wall
	}
	g.SetValue(xPos+xDelta, yPos+yDelta, this)
	g.SetValue(xPos, yPos, thing)
	return behind
}

func sumGpsCoords(g *Grid) int {
	sum := 0
	for y := range g.grid {
		for x := range g.grid[y] {
			if g.GetValue(x, y) == Box {
				sum += y*100 + x
			}
		}
	}

	return sum
}

func main() {
	gridStr, movesStr, err := readFile("../aoc24-15-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := readGrid(gridStr)

	for _, m := range movesStr {
		grid = move(grid, byte(m))
	}

	fmt.Println(sumGpsCoords(grid))
}

func readGrid(input []string) *Grid {
	grid := &Grid{}

	for _, line := range input {
		grid.grid = append(grid.grid, []byte(line))
	}

	grid.xLen = len(grid.grid[0])
	grid.yLen = len(grid.grid)

	return grid
}

func readFile(filename string) (gridStr []string, movesStr string, err error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", fmt.Errorf("failed to open file: %w", err)
	}

	fileLines := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")

	i := 0
	for _, line := range fileLines {
		if line == "" {
			break
		}

		gridStr = append(gridStr, line)
		i++
	}

	movesStr = strings.Join(fileLines[i+1:], "")

	return gridStr, movesStr, nil
}
