package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func sumNumberOfMiddlePageOfCorrectUpdates(input string) int {
	// Parse input
	weAreOnTheRulesSection := true

	rules := map[int][]int{}
	updates := [][]int{}

	// parse
	for _, line := range strings.Split(input, "\n") {
		if weAreOnTheRulesSection {
			if line == "" {
				weAreOnTheRulesSection = false
				continue
			}
			rulePages := strings.Split(line, "|")
			keyPage, _ := strconv.Atoi(rulePages[0])
			pageAfter, _ := strconv.Atoi(rulePages[1])
			rules[keyPage] = append(rules[keyPage], pageAfter)
		} else { // we are on the updates section
			updatePageStrings := strings.Split(line, ",")
			updatePages := []int{}
			for _, page := range updatePageStrings {
				pageInt, _ := strconv.Atoi(page)
				updatePages = append(updatePages, pageInt)
			}
			updates = append(updates, updatePages)
		}
	}

	// check and sum
	sum := 0
	for _, update := range updates {
		if isUpdateCorrect(rules, update) {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func isUpdateCorrect(rules map[int][]int, update []int) bool {
	updateLength := len(update)
	// just because advent of code says we need to sum the middle page numbers
	if updateLength%2 == 0 || updateLength < 3 {
		return false
	}
	for i := 0; i < updateLength-1; i++ {
		for j := i + 1; j < updateLength; j++ {
			if !isPageCorrect(rules, update[i], update[j]) {
				return false
			}
		}
	}
	return true
}

func isPageCorrect(rules map[int][]int, page, pageAfter int) bool {
	if _, ok := rules[page]; !ok {
		return false
	}
	return slices.Contains(rules[page], pageAfter)
}

func main() {
	// read values from aoc24-2-input.txt file
	input, err := readFile("../aoc24-5-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(sumNumberOfMiddlePageOfCorrectUpdates(input))
}

func readFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(file), nil
}
