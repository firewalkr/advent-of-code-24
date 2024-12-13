package main

import (
	"fmt"
	"os"
	"strings"
)

// Pos represents a position in the grid, it is 0-indexed
type Pos struct {
	X int
	Y int
}

type Grid struct {
	grid          [][]byte
	xLen          int
	yLen          int
	pastStatuses  map[Status]struct{}
	pastPositions map[Pos]struct{}
}

var errOutOfBounds = fmt.Errorf("out of bounds")
var errObstacle = fmt.Errorf("obstacle")
var errStuckInLoop = fmt.Errorf("stuck in loop")

type Node byte

const (
	up       Node = '^'
	down     Node = 'v'
	left     Node = '<'
	right    Node = '>'
	empty    Node = '.'
	obstacle Node = '#'
)

var nextDir = map[Node]Node{
	up:    right,
	right: down,
	down:  left,
	left:  up,
}

type Status struct {
	Pos Pos
	Dir Node
}

func (s *Status) String() string {
	return fmt.Sprintf("(%d, %d) %c", s.Pos.X, s.Pos.Y, s.Dir)
}

func NewGrid(input string) *Grid {
	gridStr := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]byte, len(gridStr))
	for i, row := range gridStr {
		grid[i] = []byte(row)
	}

	res := &Grid{
		grid:          grid,
		xLen:          len(grid[0]),
		yLen:          len(grid),
		pastStatuses:  map[Status]struct{}{},
		pastPositions: map[Pos]struct{}{},
	}

	s, _ := res.getCurrentStatus()

	if s != nil {
		res.pastStatuses[*s] = struct{}{}
		res.pastPositions[s.Pos] = struct{}{}
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

func (g *Grid) getValAt(pos Pos) (Node, error) {
	if g.isOutOfBounds(pos) {
		return 0, errOutOfBounds
	}
	return Node(g.grid[pos.Y][pos.X]), nil
}

func (g *Grid) setValAt(pos Pos, val Node) error {
	if g.isOutOfBounds(pos) {
		return errOutOfBounds
	}

	g.grid[pos.Y][pos.X] = byte(val)

	return nil
}

func (g *Grid) isOutOfBounds(pos Pos) bool {
	return pos.X < 0 || pos.X >= g.xLen || pos.Y < 0 || pos.Y >= g.yLen
}

func (g *Grid) getCurrentPosition() (*Pos, error) {
	for y, row := range g.grid {
		for x, val := range row {
			orientation := Node(val)
			if orientation == up || orientation == down || orientation == left || orientation == right {
				return &Pos{X: x, Y: y}, nil
			}
		}
	}

	return nil, errOutOfBounds
}

func (g *Grid) getMaybeNextPosition() (*Pos, error) {
	s, err := g.getCurrentStatus()
	if err != nil {
		return nil, err
	}

	p := s.Pos

	switch s.Dir {
	case up:
		p.Y--
	case down:
		p.Y++
	case left:
		p.X--
	case right:
		p.X++
	}

	if g.isOutOfBounds(p) {
		return nil, errOutOfBounds
	}

	if Node(g.grid[p.Y][p.X]) == obstacle {
		return nil, errObstacle
	}

	maybeNewStatus := Status{
		Pos: p,
		Dir: s.Dir,
	}

	if _, ok := g.pastStatuses[maybeNewStatus]; ok {
		return nil, errStuckInLoop
	}

	return &p, nil
}

func (g *Grid) getCurrentStatus() (*Status, error) {
	pos, err := g.getCurrentPosition()
	if err != nil {
		return nil, err
	}

	val, err := g.getValAt(*pos)
	if err != nil {
		return nil, err
	}

	return &Status{
		Pos: *pos,
		Dir: Node(val),
	}, nil
}

func (g *Grid) move() (*Status, error) {
	s, err := g.getCurrentStatus()
	if err != nil {
		return nil, err
	}

	fmt.Println("current status", s)

	tries := 0

	var newPos Pos
	dir := s.Dir
	for {
		tries++
		if tries > 4 {
			return nil, errStuckInLoop
		}

		p, err := g.getMaybeNextPosition()
		if err == errObstacle {
			// Turn 90 degrees clockwise
			dir = nextDir[dir]
			g.setValAt(s.Pos, dir)
		} else if err == errOutOfBounds {
			g.setValAt(s.Pos, empty)
			return nil, errOutOfBounds
		} else if err != nil {
			return nil, err
		} else {
			newPos = *p
			break
		}
	}

	// Update current position
	g.setValAt(s.Pos, empty)
	g.setValAt(newPos, dir)

	newStatus := Status{
		Pos: newPos,
		Dir: dir,
	}

	g.pastStatuses[newStatus] = struct{}{}
	g.pastPositions[newStatus.Pos] = struct{}{}

	return &newStatus, nil
}

func run(grid *Grid) error {
	for {
		// fmt.Printf("%s\n\n\n", grid)
		_, err := grid.move()
		if err == errOutOfBounds {
			fmt.Println("Out of bounds")
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// read values from aoc24-2-input.txt file
	input, err := readFile("../aoc24-6-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := NewGrid(input)

	err = run(grid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(grid)

	fmt.Println(len(grid.pastPositions))
}

func readFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(file), nil
}
