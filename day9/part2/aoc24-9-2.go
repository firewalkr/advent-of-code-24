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

	currentFileID := 0
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			currentFileID = disk[i]
			break
		}
	}

	currentRightPos, fileSize := previousFileBlock(disk, currentRightPos, currentFileID)
	for ; currentFileID > 0 && currentRightPos >= 0; currentFileID-- {
		currentLeftPos = nextEmptyBlock(disk, currentLeftPos, fileSize)

		if currentLeftPos == -1 || currentLeftPos > currentRightPos {
			currentLeftPos = 0
			currentRightPos, fileSize = previousFileBlock(disk, currentRightPos, currentFileID-1)
			continue
		}

		for i := 0; i < fileSize; i++ {
			disk[currentLeftPos+i] = currentFileID
			disk[currentRightPos+i] = -1
		}

		currentLeftPos = 0
		currentRightPos, fileSize = previousFileBlock(disk, currentRightPos, currentFileID-1)
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

func nextEmptyBlock(disk []int, pos int, minSize int) int {
	emptyBlockSize := 0
	for i := pos; i < len(disk); i++ {
		if disk[i] == -1 {
			emptyBlockSize++
		} else {
			emptyBlockSize = 0
		}
		if emptyBlockSize >= minSize {
			return i - minSize + 1
		}
	}

	return -1
}

func previousFileBlock(disk []int, pos, fileID int) (newPos, fileSize int) {
	newPos = pos
	for ; newPos >= 0 && disk[newPos] != fileID; newPos-- {
	}

	if newPos >= 0 {
		fileSize = 0
		for ; newPos >= 0 && disk[newPos] == fileID; newPos-- {
			fileSize++
		}

		newPos++
		return
	}

	return
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
