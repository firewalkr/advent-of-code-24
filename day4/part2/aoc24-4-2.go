package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func countCrossedMAS(input [][]byte) int {
	MAS := []byte("MAS")
	SAM := []byte("SAM")
	length := 3

	count := 0
	for row := 1; row < len(input)-1; row++ {
		for col := 1; col < len(input[row])-1; col++ {
			current := input[row][col]
			if current == 'A' {
				dd := getDiagonalDownByteSlice(input, row-1, col-1, length)
				du := getDiagonalUpByteSlice(input, row+1, col-1, length)
				if (bytes.Equal(dd, MAS) || bytes.Equal(dd, SAM)) && (bytes.Equal(du, MAS) || bytes.Equal(du, SAM)) {
					count++
				}
			}
		}
	}

	return count
}

func getDiagonalDownByteSlice(input [][]byte, row, col, length int) []byte {
	diagonal := make([]byte, length)
	for i := 0; i < length; i++ {
		diagonal[i] = input[row+i][col+i]
	}
	return diagonal
}

func getDiagonalUpByteSlice(input [][]byte, row, col, length int) []byte {
	diagonal := make([]byte, length)
	for i := 0; i < length; i++ {
		diagonal[i] = input[row-i][col+i]
	}
	return diagonal
}

func main() {
	// read values from aoc24-2-input.txt file
	input, err := readFile("../aoc24-4-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(countCrossedMAS(input))
}

func readFile(filename string) ([][]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	res := inputAsArrayOfByteArrays(string(file))
	return res, nil
}

func inputAsArrayOfByteArrays(input string) [][]byte {
	lines := strings.Split(input, "\n")
	var result [][]byte
	for _, line := range lines {
		if len(line) > 0 {
			result = append(result, []byte(line))
		}
	}
	return result
}
