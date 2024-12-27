package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const gridSizeX = 101
const gridSizeY = 103

var rgxBot = regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

type Bot struct {
	xPos, yPos   int
	xStep, yStep int
}

func moveBot(b *Bot, numSteps, gridSizeX, gridSizeY int) {
	b.xPos = trimPos(b.xPos+numSteps*b.xStep, gridSizeX) % gridSizeX
	b.yPos = trimPos(b.yPos+numSteps*b.yStep, gridSizeY) % gridSizeY
}

func trimPos(pos, size int) int {
	if pos < 0 {
		pos = -pos % size
		return size - pos
	}
	return pos % size
}

type Quadrants struct {
	topLeft, topRight, bottomLeft, bottomRight int
}

func (q *Quadrants) safetyFactor() int {
	return q.topLeft * q.bottomLeft * q.topRight * q.bottomRight
}

func getQuadrants(bots []*Bot, gridSizeX, gridSizeY int) Quadrants {
	q := Quadrants{}

	midX := gridSizeX / 2
	midY := gridSizeY / 2

	for _, b := range bots {
		switch {
		case b.xPos < midX && b.yPos < midY:
			q.topLeft++
		case b.xPos > midX && b.yPos < midY:
			q.topRight++
		case b.xPos < midX && b.yPos > midY:
			q.bottomLeft++
		case b.xPos > midX && b.yPos > midY:
			q.bottomRight++
		}
	}

	return q
}

func calcSafetyFactorAfter(bots []*Bot, numSteps, gridSizeX, gridSizeY int) int {
	for _, b := range bots {
		moveBot(b, numSteps, gridSizeX, gridSizeY)
	}

	q := getQuadrants(bots, gridSizeX, gridSizeY)

	return q.safetyFactor()
}

func main() {
	input, err := readFile("../aoc24-14-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bots := readBots(input)

	// after testing with the below approaches I found the tree
	// pattern at 7790 steps

	for _, b := range bots {
		moveBot(b, 7790, gridSizeX, gridSizeY)
	}

	printGrid(bots, gridSizeX, gridSizeY)

	// move bots a preliminary 7775 steps
	// since I know the answer is between 7500 and 10K
	// by inputting test answers on the AoC form
	// and I saw the tree show up close to 7800 steps

	// for _, b := range bots {
	// 	moveBot(b, 7775, gridSizeX, gridSizeY)
	// }

	// i := 0
	// for ; ; i++ {
	// 	// if isMaybeChristmasTreePattern(bots, gridSizeX, gridSizeY) {
	// 	printGrid(bots, gridSizeX, gridSizeY)
	// 	// fmt.Println("Found possible Christmas tree pattern after", i, "steps")
	// 	// wait for user to press return
	// 	// fmt.Scanln()
	// 	// }

	// 	for _, b := range bots {
	// 		moveBot(b, 1, gridSizeX, gridSizeY)
	// 	}

	// 	// if i%10000 == 0 {
	// 	fmt.Println("After", i, "steps")
	// 	// }

	// 	// fmt.Scanln()
	// 	time.Sleep(200 * time.Millisecond)
	// }

	// i := 0
	// for ; ; i++ {
	// 	printGrid(bots, gridSizeX, gridSizeY)
	// 	fmt.Println("After", i, "steps")

	// 	// move every bot
	// 	for _, b := range bots {
	// 		moveBot(b, 1, gridSizeX, gridSizeY)
	// 	}

	// 	fmt.Scanln()
	// }
}

func readBots(fileContent string) []*Bot {
	bots := []*Bot{}

	for _, line := range strings.Split(strings.TrimSpace(fileContent), "\n") {
		matches := rgxBot.FindStringSubmatch(line)
		xPos, _ := strconv.Atoi(matches[1])
		yPos, _ := strconv.Atoi(matches[2])
		xStep, _ := strconv.Atoi(matches[3])
		yStep, _ := strconv.Atoi(matches[4])
		bots = append(bots, &Bot{xPos, yPos, xStep, yStep})
	}

	return bots
}

// this ended up not being used.
// I decided to try and binary search the results on the AoC form
// and after trying 10K and 5K and 7.5K I knew the answer was between 7.5K and 10K.
func isMaybeChristmasTreePattern(bots []*Bot, gridSizeX, gridSizeY int) bool {
	//grid := getGrid(bots, gridSizeX, gridSizeY)

	// FIRST APPROACH
	//
	// // let's assume a typical Christmas tree pattern, i.e. something beginning like this:
	// // ....X....
	// // ...X.X...
	// // ..X...X..
	// // .X.....X.
	// // ..etc...
	// // and look at the first 5 rows

	// // first row
	// if grid[0][gridSizeX/2] != 'X' {
	// 	return false
	// }

	// for y := 1; y < 5; y++ {
	// 	if grid[y][gridSizeX/2] != '.' {
	// 		return false
	// 	}

	// 	if grid[y][gridSizeX/2-1] != 'X' || grid[y][gridSizeX/2+1] != 'X' {
	// 		return false
	// 	}
	// }

	// SECOND APPROACH
	//
	// // let's assume the tree is symmetrical
	// // and thus, search for a pattern where
	// // the left half of the grid is the same as the right half

	// for y := 0; y < gridSizeY; y++ {
	// 	for x := 0; x < gridSizeX/2; x++ {
	// 		if grid[y][x] != grid[y][gridSizeX-x-1] {
	// 			return false
	// 		}
	// 	}
	// }

	return true
}

func printGrid(bots []*Bot, gridSizeX, gridSizeY int) {
	// Clear the screen and move the cursor to the top-left
	fmt.Print("\033[H\033[2J")

	grid := getGrid(bots, gridSizeX, gridSizeY)
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func getGrid(bots []*Bot, gridSizeX, gridSizeY int) [][]byte {
	grid := make([][]byte, gridSizeY)
	for y := range grid {
		grid[y] = make([]byte, gridSizeX)
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}

	for _, b := range bots {
		grid[b.yPos][b.xPos] = 'X'
	}

	return grid
}

func readFile(filename string) (string, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return strings.TrimSpace(string(fileBytes)), nil
}
