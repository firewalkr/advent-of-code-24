package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(input string) bool {
	// Parse input
	levels, err := parseInput(input)
	if err != nil {
		panic(err)
	}

	if len(levels) < 2 {
		return false
	}

	if levels[0] < levels[1] {
		return checkIncrease(levels)
	} else if levels[0] > levels[1] {
		return checkDecrease(levels)
	}

	return false
}

func checkDecrease(levels []int) bool {
	previous := levels[0]
	// Check if levels are safe
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - previous
		if diff < -3 || diff > -1 {
			return false
		}
		previous = levels[i]
	}

	return true
}

func checkIncrease(levels []int) bool {
	previous := levels[0]
	// Check if levels are safe
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - previous
		if diff > 3 || diff < 1 {
			return false
		}
		previous = levels[i]
	}

	return true
}

func parseInput(input string) ([]int, error) {
	var err error
	strLevels := strings.Split(input, " ")
	levels := make([]int, len(strLevels))
	for i, strLevel := range strLevels {
		levels[i], err = strconv.Atoi(strLevel)
		if err != nil {
			return nil, err
		}
	}

	return levels, nil
}

func main() {
	// read values from aoc24-2-input.txt file
	listsOfNumbers, err := readFile("../aoc24-2-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	safeReportCount := 0
	for _, numbers := range listsOfNumbers {
		if isReportSafe(numbers) {
			safeReportCount++
		}
	}

	fmt.Println(safeReportCount)
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	listsOfNumbers := make([]string, 0)

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		listsOfNumbers = append(listsOfNumbers, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return listsOfNumbers, nil
}
