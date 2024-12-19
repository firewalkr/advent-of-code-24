package main

import (
	"fmt"
	"os"
)

func mapBlocks(input string) []int {
	out := []int{}
	isFiles := true
	fileID := 0
	for _, c := range input {
		num := int(c) - 48
		for i := 0; i < num; i++ {
			if isFiles {
				out = append(out, fileID)
			} else {
				out = append(out, -1)
			}
		}

		if isFiles {
			fileID++
		}
		isFiles = !isFiles
	}

	return out
}

func moveBlocks(disk []int) []int {
	currentLeftPos := 0
	currentRightPos := len(disk) - 1
	for {
		currentLeftPos := nextEmptyBlock(disk, currentLeftPos)
		currentRightPos := previousFileBlock(disk, currentRightPos)

		if currentLeftPos == -1 || currentRightPos == -1 || currentRightPos < currentLeftPos {
			break
		}

		disk[currentLeftPos] = disk[currentRightPos]
		disk[currentRightPos] = -1
		currentLeftPos++
	}

	return disk
}

func checksum(disk []int) int {
	count := 0
	for i, c := range disk {
		if c != -1 {
			count += i * c
		}
	}

	return count
}

func nextEmptyBlock(disk []int, pos int) int {
	for i := pos; i < len(disk); i++ {
		if disk[i] == -1 {
			return i
		}
	}

	return -1
}

func previousFileBlock(disk []int, pos int) int {
	for i := pos; i >= 0; i-- {
		if disk[i] != -1 {
			return i
		}
	}

	return -1
}

func doAll(input string) int {
	return checksum(moveBlocks(mapBlocks(input)))
}

func main() {
	input, err := readFile("../aoc24-9-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(doAll(input))
}

func readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(fileContent), nil
}
