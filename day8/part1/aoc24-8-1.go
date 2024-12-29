package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// 0-based
type Pos struct {
	X int
	Y int
}

func isAntinodeValid(a Pos, sizeX, sizeY int) bool {
	return a.X >= 0 && a.Y >= 0 && a.X < sizeX && a.Y < sizeY
}

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func getAntinodes(p1, p2 Pos, sizeX, sizeY int) []Pos {
	diffX := intAbs(p1.X - p2.X)
	diffY := intAbs(p1.Y - p2.Y)

	var left, right *Pos
	if p1.X-p2.X < 0 {
		left = &p1
		right = &p2
	} else {
		left = &p2
		right = &p1
	}
	antiNodeLeft := Pos{X: left.X - diffX, Y: left.Y - diffY}
	antiNodeRight := Pos{X: right.X + diffX, Y: right.Y + diffY}

	if left.Y < right.Y {
		antiNodeLeft.Y = left.Y - diffY
		antiNodeRight.Y = right.Y + diffY
	} else {
		antiNodeLeft.Y = left.Y + diffY
		antiNodeRight.Y = right.Y - diffY
	}

	results := []Pos{}
	if isAntinodeValid(antiNodeLeft, sizeX, sizeY) {
		results = append(results, antiNodeLeft)
	}
	if isAntinodeValid(antiNodeRight, sizeX, sizeY) {
		results = append(results, antiNodeRight)
	}

	return results
}

func CountAntinodes(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sizeX := len(lines[0])
	sizeY := len(lines)

	positionsByChar := map[byte][]Pos{}

	for y, line := range lines {
		for x := range len(line) {
			b := byte(line[x])
			if b != '.' {
				if _, ok := positionsByChar[b]; !ok {
					positionsByChar[b] = []Pos{}
				}
				positionsByChar[b] = append(positionsByChar[b], Pos{x, y})
			}
		}
	}

	for b := range positionsByChar {
		if len(positionsByChar[b]) < 2 {
			delete(positionsByChar, b)
		}
	}

	positionsWithAntinodes := map[Pos]struct{}{}

	for _, positions := range positionsByChar {
		for i := 0; i < len(positions)-1; i++ {
			for j := i + 1; j < len(positions); j++ {
				antinodes := getAntinodes(positions[i], positions[j], sizeX, sizeY)
				for _, a := range antinodes {
					positionsWithAntinodes[a] = struct{}{}
				}
			}
		}
	}

	// fmt.Println(positionsWithAntinodes)
	// printResult(input, positionsWithAntinodes)

	return len(positionsWithAntinodes)
}

func main() {
	input, err := readFile("../aoc24-8-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(CountAntinodes(input))
}

func printResult(input string, positionsWithAntiNodes map[Pos]struct{}) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, line := range lines {
		bs := []byte(line)
		for x := range len(line) {
			if _, hasAntinode := positionsWithAntiNodes[Pos{x, y}]; hasAntinode && line[x] == '.' {
				bs[x] = '#'
			}
		}
		fmt.Println(string(bs))
	}
}

func readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(fileContent), nil
}
