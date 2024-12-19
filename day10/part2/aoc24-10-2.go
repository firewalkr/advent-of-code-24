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

// actually wrote this for part 1, coz distractedly I thought it wanted the total paths per trailhead
func countPathsPerTrail(terrain [][]int8, xStart, yStart int, expected int8, numReached int) int {
	xSize, ySize := dims(terrain)

	if xStart < 0 || yStart < 0 || xStart >= xSize || yStart >= ySize {
		return 0
	}

	if terrain[yStart][xStart] != expected {
		return 0
	}

	if expected == 9 {
		return numReached + 1
	}

	leftToRight := countPathsPerTrail(terrain, xStart+1, yStart, expected+1, numReached)
	rightToLeft := countPathsPerTrail(terrain, xStart-1, yStart, expected+1, numReached)
	topToBottom := countPathsPerTrail(terrain, xStart, yStart+1, expected+1, numReached)
	bottomToTop := countPathsPerTrail(terrain, xStart, yStart-1, expected+1, numReached)

	return leftToRight + rightToLeft + topToBottom + bottomToTop
}

type Pos struct {
	x, y int
}

func countAllTrailheads(terrain [][]int8) int {
	xSize, ySize := dims(terrain)

	if xSize == 0 || ySize == 0 {
		return 0
	}

	rating := 0
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if terrain[y][x] == 0 {
				rating += countPathsPerTrail(terrain, x, y, 0, 0)
			}
		}
	}

	return rating
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

	fmt.Println(countAllTrailheads(stringToTerrain(input)))
}

func readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(fileContent), nil
}
