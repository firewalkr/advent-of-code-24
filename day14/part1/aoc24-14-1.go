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

	fmt.Println(calcSafetyFactorAfter(bots, 100, gridSizeX, gridSizeY))
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

func readFile(filename string) (string, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return strings.TrimSpace(string(fileBytes)), nil
}
