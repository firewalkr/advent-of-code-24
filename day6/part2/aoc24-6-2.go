package main

import (
	"fmt"
	"os"
	"strings"
)

type Node byte

type Grid struct {
	grid [][]Node
	xLen int
	yLen int
}

var errStuckInLoop = fmt.Errorf("stuck in loop")

const (
	up          Node = '^'
	down        Node = 'v'
	left        Node = '<'
	right       Node = '>'
	empty       Node = '.'
	obstacle    Node = '#'
	outOfBounds Node = 0
)

func NewGrid(input string) *Grid {
	gridStr := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]Node, len(gridStr))
	for i, row := range gridStr {
		grid[i] = []Node(row)
	}

	res := &Grid{
		grid: grid,
		xLen: len(grid[0]),
		yLen: len(grid),
	}

	return res
}

func (g *Grid) String() string {
	var sb strings.Builder
	for _, row := range g.grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}

	return strings.TrimSpace(sb.String())
}

func (g *Grid) getValAt(x, y int) Node {
	if g.isOutOfBounds(x, y) {
		return outOfBounds
	}

	return Node(g.grid[y][x])
}

func (g *Grid) setValAt(x, y int, val Node) {
	g.grid[y][x] = val
}

func (g *Grid) isOutOfBounds(x, y int) bool {
	return x < 0 || x >= g.xLen || y < 0 || y >= g.yLen
}

func (g *Grid) getBotPosition() (int, int) {
	for y, row := range g.grid {
		for x, val := range row {
			orientation := Node(val)
			if orientation == up || orientation == down || orientation == left || orientation == right {
				return x, y
			}
		}
	}

	return -1, -1
}

type Status struct {
	x, y, deltaX, deltaY int
}

type Pos struct {
	x, y int
}

func run(grid *Grid, botX, botY int, direction Node) (map[Status]struct{}, error) {
	deltaX, deltaY := getDeltas(direction)

	wentBy := map[Status]struct{}{}

	for {
		// fmt.Print("\033[H\033[2J")
		// grid.setValAt(botX, botY, '@')
		// fmt.Print(grid)
		// grid.setValAt(botX, botY, empty)
		// time.Sleep(200 * time.Millisecond)

		s := Status{botX, botY, deltaX, deltaY}
		if _, been := wentBy[s]; been {
			return wentBy, errStuckInLoop
		}
		wentBy[s] = struct{}{}

		newX := botX + deltaX
		newY := botY + deltaY
		node := grid.getValAt(newX, newY)
		if node == outOfBounds {
			break
		}
		if node == obstacle {
			deltaX, deltaY = rotate90CW(deltaX, deltaY)
			continue
		}
		botX = newX
		botY = newY
	}

	return wentBy, nil
}

func getDeltas(direction Node) (int, int) {
	deltaX, deltaY := 0, 0

	if direction == up {
		deltaY = -1
	} else if direction == right {
		deltaX = 1
	} else if direction == down {
		deltaY = 1
	} else if direction == left {
		deltaX = -1
	}

	return deltaX, deltaY
}

// assumes consistent delta values
func rotate90CW(curDeltaX, curDeltaY int) (int, int) {
	if curDeltaY == -1 {
		return 1, 0
	} else if curDeltaX == 1 {
		return 0, 1
	} else if curDeltaY == 1 {
		return -1, 0
	} else { // if curDeltaX == -1
		return 0, -1
	}
}

func main() {
	// read values from aoc24-2-input.txt file
	input, err := readFile("../aoc24-6-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := NewGrid(input)

	botX, botY := grid.getBotPosition()
	directionNode := grid.getValAt(botX, botY)
	grid.setValAt(botX, botY, empty)

	// do a first run to check where the bot passes by without extra obstacles.
	// we know we need to place the test obstacle only in these locations.
	noObstacleStatuses, err := run(grid, botX, botY, directionNode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	positionsFirstRun := map[Pos]struct{}{}
	for s := range noObstacleStatuses {
		positionsFirstRun[Pos{s.x, s.y}] = struct{}{}
	}

	fmt.Printf("grid dimensions: %d x %d\n", grid.xLen, grid.yLen)

	var countPositionsThatLeadToLoop int

	for x := 0; x < grid.xLen; x++ {
		for y := 0; y < grid.yLen; y++ {
			if _, goesByPosition := positionsFirstRun[Pos{x, y}]; !goesByPosition {
				continue
			}

			v := grid.getValAt(x, y)
			if v == empty {
				// fmt.Printf("Trying %d, %d\n", x, y)
				// clone grid.
				grid.setValAt(x, y, obstacle)

				_, err := run(grid, botX, botY, directionNode)
				if err == errStuckInLoop {
					countPositionsThatLeadToLoop++
				}

				grid.setValAt(x, y, empty)
			}
		}
	}

	fmt.Println(countPositionsThatLeadToLoop)
}

func readFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(file), nil
}
