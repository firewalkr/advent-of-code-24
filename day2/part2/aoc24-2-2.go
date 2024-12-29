package main

import (
	"bufio"
	"fmt"

	// "log/slog"
	"math"
	"os"
	"strconv"
	"strings"
)

func isReportSafeWithAtMostOneFlaw_Ugly(input string) bool {
	// Parse input
	levels, err := parseInput(input)
	if err != nil {
		panic(err)
	}

	if len(levels) < 2 {
		return false
	}

	if checkSequenceStrict(levels) {
		return true
	}

	// brute force time
	numSuccesses := 0
	for i := 0; i < len(levels); i++ {
		levelsWithoutCurrent := deleteAtIndex(levels, i)
		// slog.Info("Here we go", "levels", levels, "levelsWithoutCurrent", levelsWithoutCurrent)
		if !checkSequenceStrict(levelsWithoutCurrent) {
			// slog.Info("Failed")
		} else {
			numSuccesses++
			// slog.Info("Success")
		}
	}

	return numSuccesses > 0
}

func deleteAtIndex(slice []int, index int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	for i, val := range slice {
		if i == index {
			continue
		}
		newSlice = append(newSlice, val)
	}
	return newSlice
}

func checkSequenceStrict(levels []int) bool {
	previous := levels[0]
	increasing := levels[1] > previous
	// Check if levels are safe
	for i := 1; i < len(levels); i++ {
		absDiff := int(math.Abs(float64(levels[i] - previous)))
		nowIncreasing := levels[i] > previous
		if increasing != nowIncreasing {
			return false
		}
		// slog.Info(fmt.Sprintf("current: %d, previous: %d, absDiff: %d", levels[i], previous, absDiff))

		if absDiff > 3 || absDiff == 0 {
			return false
		}

		previous = levels[i]
	}

	return true
}

// below was me trying to be smart. there's almost certainly
// a better way than isReportSafeWithAtMostOneFlaw_Ugly

// func isReportSafeWithAtMostOneFlaw(input string) bool {
// 	// Parse input
// 	levels, err := parseInput(input)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if len(levels) < 2 {
// 		return false
// 	}
// 	slog.Info("Here we go", "levels", levels)
// 	return checkSequence(levels)
// }

// func checkSequence(levels []int) bool {
// 	previous := levels[0]
// 	flawCount := 0
// 	prevDiff := levels[1] - previous
// 	// Check if levels are safe
// 	for i := 1; i < len(levels); i++ {
// 		diff := levels[i] - previous
// 		absDiff := int(math.Abs(float64(diff)))
// 		slog.Info(fmt.Sprintf("current: %d, previous: %d, diff: %d, absDiff: %d, prevDiff: %d", levels[i], previous, diff, absDiff, prevDiff))
// 		if absDiff > 3 || absDiff < 1 || (prevDiff*diff) < 0 {
// 			if flawCount == 1 {
// 				return false
// 			}
// 			flawCount++
// 			continue
// 		}
// 		prevDiff = diff
// 		previous = levels[i]
// 	}

// 	return true
// }

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
		if isReportSafeWithAtMostOneFlaw_Ugly(numbers) {
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
