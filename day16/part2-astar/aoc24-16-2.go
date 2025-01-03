package main

import (
	"fmt"
	"math"
	"os"
	"runtime/pprof"
	"strings"
)

// Implemented A* like a boss (not really), via the Wikipedia pseudocode, to
// see what real code looks and performs like. This could still use a better
// structure for the openSet, since I'm always finding the minimum fScore.
//
// The initial implementation didn't get the nodes from _all_ the best paths,
// of course, so I then did a (very) quick fix on the A* implementation, i.e.
// I changed `cameFrom` to store a slice of nodes instead of just one.
//
// Claude Sonnet did the changes on the path reconstruction (after I changed the
// function signature, but Sonnet did it pretty much like an actual boss).
//
// Anyway, I clearly had forgotten how this worked (if I ever knew but I'm pretty
// sure I touched it in college) and now I see why it's so good. The advantage
// of having done my own poor man's version is that I now understand A* better
// AND I was also able to do the "get all nodes from all best paths" bit in just
// a couple of minutes, since I knew exactly how to modify A*.

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

type Pos struct {
	x, y int
}

type Traversal struct {
	pos       Pos
	direction Direction
}

func scoreMapGet(gScore map[Traversal]int, t Traversal) int {
	if score, seen := gScore[t]; seen {
		return score
	}

	return math.MaxInt
}

// here, we only care about the positions traversed, not the directions
func reconstructPath(cameFrom map[Traversal][]Traversal, current Traversal) []Pos {
	totalPath := []Pos{current.pos}

	currents := []Traversal{current}

	for {
		nextCurrents := []Traversal{}

		for _, c := range currents {
			if previous, seen := cameFrom[c]; seen {
				for _, p := range previous {
					totalPath = append(totalPath, p.pos)
					nextCurrents = append(nextCurrents, p)
				}
			}
		}

		if len(nextCurrents) == 0 {
			break
		}

		currents = nextCurrents
	}

	return uniq(totalPath)
}

func uniq(positions []Pos) []Pos {
	seen := map[Pos]struct{}{}
	uniq := []Pos{}

	for _, p := range positions {
		if _, ok := seen[p]; !ok {
			uniq = append(uniq, p)
			seen[p] = struct{}{}
		}
	}

	return uniq
}

func aStar(
	g *Grid,
	heuristic func(t Traversal) int,
) ([]Pos, int) {
	start := Pos{g.startX, g.startY}
	end := Pos{g.endX, g.endY}

	startTraversal := Traversal{start, East}

	openSet := map[Traversal]struct{}{startTraversal: {}}

	cameFrom := map[Traversal][]Traversal{}

	gScore := map[Traversal]int{}
	gScore[startTraversal] = 0

	fScore := map[Traversal]int{}
	fScore[startTraversal] = heuristic(startTraversal)

	for len(openSet) > 0 {
		currentNode := getMinFScoreNode(openSet, fScore)

		if currentNode.pos == end {
			return reconstructPath(cameFrom, currentNode), scoreMapGet(gScore, currentNode)
		}

		delete(openSet, currentNode)

		for _, neighbour := range getNeighbours(g, currentNode) {
			tentativeGScore := scoreMapGet(gScore, currentNode) + distanceToNeighbour(currentNode, neighbour)

			if tentativeGScore == scoreMapGet(gScore, neighbour) {
				cameFrom[neighbour] = append(cameFrom[neighbour], currentNode)
			}

			if tentativeGScore < scoreMapGet(gScore, neighbour) {
				cameFrom[neighbour] = append(cameFrom[neighbour], currentNode)
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = gScore[neighbour] + heuristic(neighbour)

				if _, seen := openSet[neighbour]; !seen {
					openSet[neighbour] = struct{}{}
				}
			}
		}
	}

	return []Pos{}, math.MaxInt
}

// rotations cost 1000, movement costs 1
func distanceToNeighbour(start, neighbour Traversal) int {
	if start.direction != neighbour.direction {
		return 1000
	}

	return 1
}

func getMinFScoreNode(openSet map[Traversal]struct{}, fScore map[Traversal]int) Traversal {
	minScore := math.MaxInt
	var minTraversal Traversal

	for t := range openSet {
		if fScore[t] < minScore {
			minScore = fScore[t]
			minTraversal = t
		}
	}

	return minTraversal
}

func getNeighbours(g *Grid, t Traversal) []Traversal {
	neighbours := []Traversal{}

	deltaX, deltaY := t.direction.deltas()
	if g.GetValue(t.pos.x+deltaX, t.pos.y+deltaY) != Wall {
		neighbours = append(neighbours, Traversal{Pos{t.pos.x + deltaX, t.pos.y + deltaY}, t.direction})
	}

	direction90CW := t.direction.rotate90CW()
	deltaX, deltaY = direction90CW.deltas()
	if g.GetValue(t.pos.x+deltaX, t.pos.y+deltaY) != Wall {
		neighbours = append(neighbours, Traversal{Pos{t.pos.x, t.pos.y}, direction90CW})
	}

	direction90CCW := t.direction.rotate90CCW()
	deltaX, deltaY = direction90CCW.deltas()
	if g.GetValue(t.pos.x+deltaX, t.pos.y+deltaY) != Wall {
		neighbours = append(neighbours, Traversal{Pos{t.pos.x, t.pos.y}, direction90CCW})
	}

	return neighbours
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

	path, score := aStar(grid, func(t Traversal) int {
		return abs(t.pos.x-grid.endX) + abs(t.pos.y-grid.endY) // FIXME: Manhattan distance is not necessarily good here, because of rotation costs
	})

	fmt.Println(len(path))
	fmt.Println(score)

	posMap := make(map[Pos]struct{})
	for _, pos := range path {
		posMap[pos] = struct{}{}
	}

	fmt.Println(grid.PrintWithTilesInBestPaths(posMap))
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

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
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
