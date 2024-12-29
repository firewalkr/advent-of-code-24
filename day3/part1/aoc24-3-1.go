package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "log"
)

func addMuls(input string) int {
	// Parse input
	regexp := regexp.MustCompile(`mul\((-?[0-9]+),(-?[0-9]+)\)|do\(\)|don't\(\)`)

	matches := regexp.FindAllStringSubmatch(input, -1)

	// fmt.Println("matches", "matches", matches)

	sum := 0
	multiplicationActive := true
	for _, match := range matches {
		if match[0] == "do()" {
			multiplicationActive = true
			continue
		} else if match[0] == "don't()" {
			multiplicationActive = false
			continue
		}

		if multiplicationActive {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			// log.Println("match", "num1", num1, "num2", num2)

			sum += num1 * num2
		}
	}

	return sum
}

func main() {
	// read values from aoc24-2-input.txt file
	input, err := readFile("../aoc24-3-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(addMuls(input))
}

func readFile(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(file), nil
}
