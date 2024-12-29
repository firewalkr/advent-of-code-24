package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func countXMAS(input [][]byte) int {
	XMAS := []byte("XMAS")
	SAMX := []byte("SAMX")
	lenXMAS := len(XMAS)

	count := 0
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if col <= len(input[row])-lenXMAS {
				horiz := input[row][col : col+lenXMAS]
				// fmt.Printf("comparing %s\n", horiz)
				if bytes.Equal(horiz, XMAS) {
					count++
				} else if bytes.Equal(horiz, SAMX) {
					count++
				}
			}
			if row <= len(input)-lenXMAS {
				vert := getVerticalByteSlice(input, row, col, lenXMAS)
				// fmt.Printf("comparing vertical %s\n", vert)
				if bytes.Equal(vert, XMAS) {
					count++
				} else if bytes.Equal(vert, SAMX) {
					count++
				}
			}
			if row <= len(input)-lenXMAS && col <= len(input[row])-lenXMAS {
				diagDown := getDiagonalDownByteSlice(input, row, col, lenXMAS)
				// fmt.Printf("comparing diagonal down %s\n", diagDown)
				if bytes.Equal(diagDown, XMAS) {
					count++
				} else if bytes.Equal(diagDown, SAMX) {
					count++
				}
			}
			if row >= lenXMAS-1 && col <= len(input[row])-lenXMAS {
				diagUp := getDiagonalUpByteSlice(input, row, col, lenXMAS)
				// fmt.Printf("comparing diagonal up %s\n", diagUp)
				if bytes.Equal(diagUp, XMAS) {
					count++
				} else if bytes.Equal(diagUp, SAMX) {
					count++
				}
			}
		}
	}

	return count
}

func getVerticalByteSlice(input [][]byte, row, col, length int) []byte {
	vertical := make([]byte, length)
	for i := 0; i < length; i++ {
		vertical[i] = input[row+i][col]
	}
	return vertical
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

	fmt.Println(countXMAS(input))
}

func readFile(filename string) ([][]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	res := inputAsArrayOfByteArrays(string(file))
	return res, nil
}

// input text is always 1-byte chars so it's fine
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
