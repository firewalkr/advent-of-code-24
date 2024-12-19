package main

import (
	"fmt"
	"os"
	"strings"
)

func stringToTerrain(input string) [][]int8 {
	terrain := [][]int8{}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	width := len(lines[0])

	for _, line := range lines {
		line = strings.TrimSpace(line)
		row := []int8{}
		for i := 0; i < width; i++ {
			if line[i] == '.' {
				row = append(row, -1)
			} else {
				row = append(row, int8(line[i])-'0')
			}
		}

		terrain = append(terrain, row)
	}

	return terrain
}

type Pos struct {
	x, y int
}

func aggregateReachableNines(terrain [][]int8, xStart, yStart int, expected int8, reached map[Pos]struct{}) {
	xSize, ySize := dims(terrain)

	if xStart < 0 || yStart < 0 || xStart >= xSize || yStart >= ySize {
		return
	}

	if terrain[yStart][xStart] != expected {
		return
	}

	if expected == 9 {
		reached[Pos{xStart, yStart}] = struct{}{}
		return
	}

	aggregateReachableNines(terrain, xStart+1, yStart, expected+1, reached)
	aggregateReachableNines(terrain, xStart-1, yStart, expected+1, reached)
	aggregateReachableNines(terrain, xStart, yStart+1, expected+1, reached)
	aggregateReachableNines(terrain, xStart, yStart-1, expected+1, reached)
}

func scoreAllTrailheads(terrain [][]int8) int {
	xSize, ySize := dims(terrain)

	if xSize == 0 || ySize == 0 {
		return 0
	}

	score := 0
	aggregate := map[Pos]struct{}{}
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if terrain[y][x] == 0 {
				aggregateReachableNines(terrain, x, y, 0, aggregate)
				score += len(aggregate)
				aggregate = map[Pos]struct{}{}
			}
		}
	}

	return score
}

func dims(terrain [][]int8) (xSize, ySize int) {
	ySize = len(terrain)
	if ySize == 0 {
		return
	}

	xSize = len(terrain[0])
	return
}

func main() {
	input, err := readFile("../aoc24-10-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(scoreAllTrailheads(stringToTerrain(input)))
}

func readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(fileContent), nil
}
