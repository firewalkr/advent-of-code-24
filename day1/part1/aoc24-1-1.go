package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// read values from aoc24-1-input.txt file
	listOne, listTwo, err := readFile("../aoc24-1-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(sumAbsDiffs(listOne, listTwo))
}

func sumAbsDiffs(listOne []int, listTwo []int) int {
	sort.Slice(listOne, func(i, j int) bool {
		return listOne[i] < listOne[j]
	})

	sort.Slice(listTwo, func(i, j int) bool {
		return listTwo[i] < listTwo[j]
	})

	sumAbsDiffs := 0
	for i := range len(listOne) {
		diff := listOne[i] - listTwo[i]
		if diff < 0 {
			diff = -diff
		}
		sumAbsDiffs += diff
	}

	return sumAbsDiffs
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
