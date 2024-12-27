package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func maxPresses(xStep, yStep, targetX, targetY int) int {
	maxX := targetX / xStep
	maxY := targetY / yStep

	return min(maxX, maxY)
}

const costA = 3
const costB = 1

var rgxButtonA = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
var rgxButtonB = regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
var rgxPrize = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

type Machine struct {
	buttonAX, buttonAY int
	buttonBX, buttonBY int
	prizeX, prizeY     int
}

func calcCheapestPrice(m *Machine) int {
	minimum := math.MaxInt

	maxPressesA := maxPresses(m.buttonAX, m.buttonAY, m.prizeX, m.prizeY)
	for i := maxPressesA; i >= 0; i-- {
		posx := i * m.buttonAX
		posy := i * m.buttonAY

		maxPressesB := maxPresses(m.buttonBX, m.buttonBY, m.prizeX-posx, m.prizeY-posy)
		if posx+maxPressesB*m.buttonBX == m.prizeX && posy+maxPressesB*m.buttonBY == m.prizeY {
			minimum = min(minimum, i*costA+maxPressesB*costB)
		}
	}

	return minimum
}

func main() {
	machines, err := readFile("../aoc24-13-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	totalCost := 0

	for _, m := range machines {
		minimum := calcCheapestPrice(m)
		if minimum != math.MaxInt {
			totalCost += minimum
		}
	}

	fmt.Println(totalCost)
}

func readFile(filename string) ([]*Machine, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	machines := []*Machine{}
	nextMachine := &Machine{}

	for _, line := range strings.Split(string(fileContent), "\n") {
		if strings.Contains(line, "Button A") {
			buttonAX, buttonAY := parseButtonA(line)
			nextMachine.buttonAX = buttonAX
			nextMachine.buttonAY = buttonAY
		} else if strings.Contains(line, "Button B") {
			buttonBX, buttonBY := parseButtonB(line)
			nextMachine.buttonBX = buttonBX
			nextMachine.buttonBY = buttonBY
		} else if strings.Contains(line, "Prize") {
			prizeX, prizeY := parsePrize(line)
			nextMachine.prizeX = prizeX
			nextMachine.prizeY = prizeY
			machines = append(machines, nextMachine)
			nextMachine = &Machine{}
		}
	}

	return machines, nil
}

func parseButtonA(line string) (int, int) {
	matches := rgxButtonA.FindStringSubmatch(line)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return x, y
}

func parseButtonB(line string) (int, int) {
	matches := rgxButtonB.FindStringSubmatch(line)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return x, y
}

func parsePrize(line string) (int, int) {
	matches := rgxPrize.FindStringSubmatch(line)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return x, y
}
