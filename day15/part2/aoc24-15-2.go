package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const Wall = byte('#')
const BoxLeft = byte('[')
const BoxRight = byte(']')
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
	// Clear the screen and move the cursor to the top-left
	// fmt.Print("\033[H\033[2J")
	// fmt.Println(newGrid)
	// time.Sleep(25 * time.Millisecond)

	botX, botY := g.FindRobot()

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

	g.moveBot(botX, botY, xDelta, yDelta)

	return g
}

func (g *Grid) moveBot(xPos, yPos, xDelta, yDelta int) {
	next := g.GetValue(xPos+xDelta, yPos+yDelta)
	if next == Wall {
		return
	} else if next == Empty {
		g.SetValue(xPos+xDelta, yPos+yDelta, Robot)
		g.SetValue(xPos, yPos, Empty)
		return
	}
	// else it's box
	if yDelta == 0 {
		behind := g.moveBoxesHoriz(xPos+xDelta, yPos, xDelta, Robot)
		if behind != Wall {
			g.SetValue(xPos, yPos, Empty)
		}
	} else {
		behind := g.moveBoxesVert(xPos, yPos+yDelta, yDelta)
		if behind == Empty {
			g.SetValue(xPos, yPos+yDelta, Robot)
			g.SetValue(xPos, yPos, Empty)
		}
	}

}

func (g *Grid) moveBoxesHoriz(xPos, yPos, xDelta int, thing byte) (behindBoxes byte) {
	box1 := g.GetValue(xPos, yPos)
	box2 := g.GetValue(xPos+xDelta, yPos)
	next := g.GetValue(xPos+xDelta*2, yPos)
	if next == Wall {
		return Wall
	} else if next == Empty {
		g.SetValue(xPos+xDelta*2, yPos, box2)
		g.SetValue(xPos+xDelta, yPos, box1)
		g.SetValue(xPos, yPos, thing)
		return Empty
	} // shouldn't need to check for out of bounds since the grid is surrounded by walls

	// else if Box
	behind := g.moveBoxesHoriz(xPos+xDelta, yPos, xDelta, box1)
	if behind == Wall {
		return Wall
	}
	g.SetValue(xPos+xDelta, yPos, box1)
	g.SetValue(xPos+xDelta*2, yPos, box2)
	g.SetValue(xPos, yPos, thing)
	return behind
}

type Pos struct {
	x, y int
}

// [][] [][]
//  [][][]

// breadth-first-ish, left to right
func (g *Grid) moveBoxesVert(xPos, yPos, yDelta int) (behindBoxes byte) {
	queue := []Pos{}
	boxHalf := g.GetValue(xPos, yPos)
	if boxHalf == BoxLeft {
		queue = append(queue, Pos{xPos, yPos}, Pos{xPos + 1, yPos})
	} else {
		queue = append(queue, Pos{xPos - 1, yPos}, Pos{xPos, yPos})
	}

	initial := append([]Pos{}, queue...)

	previousQueueLen := 0
	for {
		nextRowSeen := map[Pos]struct{}{}

		queueLen := len(queue)
		if len(queue) == previousQueueLen { // this row's relevant nodes are all empty
			break
		}

		// keep adding all boxes in each row to the queue
		// but on each iteration look only at the current row
		for i := previousQueueLen; i < queueLen; i++ {
			pos := queue[i]
			this := g.GetValue(pos.x, pos.y)
			ahead := g.GetValue(pos.x, pos.y+yDelta)
			if ahead == Wall {
				return Wall
			}
			if this == BoxLeft {
				if ahead == BoxRight {
					queue = append(queue, g.getPosToAddIfBoxAndNotSeen(Pos{pos.x - 1, pos.y + yDelta}, nextRowSeen)...)
				}
				queue = append(queue, g.getPosToAddIfBoxAndNotSeen(Pos{pos.x, pos.y + yDelta}, nextRowSeen)...)
			} else if this == BoxRight {
				queue = append(queue, g.getPosToAddIfBoxAndNotSeen(Pos{pos.x, pos.y + yDelta}, nextRowSeen)...)
				if ahead == BoxLeft {
					queue = append(queue, g.getPosToAddIfBoxAndNotSeen(Pos{pos.x + 1, pos.y + yDelta}, nextRowSeen)...)
				}
			}
		}

		previousQueueLen = queueLen
	}

	// move elements in queue by yDelta, running through the queue in reverse order
	for i := len(queue) - 1; i >= 0; i-- {
		pos := queue[i]
		g.SetValue(pos.x, pos.y+yDelta, g.GetValue(pos.x, pos.y))
		g.SetValue(pos.x, pos.y, Empty)
	}

	// clear the initial positions
	for _, pos := range initial {
		g.SetValue(pos.x, pos.y, Empty)
	}

	return Empty
}

func (g *Grid) getPosToAddIfBoxAndNotSeen(pos Pos, seenMap map[Pos]struct{}) []Pos {
	toAdd := []Pos{}
	_, seen := seenMap[pos]
	if !seen {
		seenMap[pos] = struct{}{}
		v := g.GetValue(pos.x, pos.y)
		if v == BoxLeft || v == BoxRight {
			toAdd = append(toAdd, pos)
		}
	}

	return toAdd
}

func sumGpsCoords(g *Grid) int {
	sum := 0
	for y := range g.grid {
		for x := range g.grid[y] {
			if g.GetValue(x, y) == BoxLeft {
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

	grid := widenGrid(readGrid(gridStr))

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

func widenGrid(grid *Grid) *Grid {
	wideGrid := &Grid{
		xLen: grid.xLen * 2,
		yLen: grid.yLen,
	}

	for y := range grid.grid {
		row := bytes.NewBuffer(make([]byte, 0, 2*len(grid.grid[y])))
		for x := range grid.grid[y] {
			switch grid.grid[y][x] {
			case '#':
				row.WriteString("##")
			case 'O':
				row.WriteString("[]")
			case '.':
				row.WriteString("..")
			case '@':
				row.WriteString("@.")
			}
		}
		wideGrid.grid = append(wideGrid.grid, row.Bytes())
	}

	return wideGrid
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
