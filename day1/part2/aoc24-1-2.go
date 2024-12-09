package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read values from aoc24-1-input.txt file
	listOne, listTwo, err := readFile("../aoc24-1-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(similarity(listOne, listTwo))
}

func similarity(listOne []int, listTwo []int) int {
	m := map[int]int{}
	for _, num := range listTwo {
		m[num]++
	}

	sim := 0
	for _, num := range listOne {
		sim += num * m[num]
	}

	return sim
}

func readFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	numbers1 := make([]int, 0)
	numbers2 := make([]int, 0)

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var num1, num2 int

		n, err := fmt.Sscanf(line, "%d%d", &num1, &num2)
		if err != nil || n != 2 {
			if err == nil {
				err = fmt.Errorf("expected 2 numbers, got %d", n)
			}
			return nil, nil, fmt.Errorf("line %d: %w", lineNum, err)
		}

		numbers1 = append(numbers1, num1)
		numbers2 = append(numbers2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %w", err)
	}

	return numbers1, numbers2, nil
}
