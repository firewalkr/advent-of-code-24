package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func stonesToString(stones []int) string {
	out := strings.Builder{}
	for _, s := range stones {
		out.WriteString(fmt.Sprintf("%d ", s))
	}

	return strings.TrimSpace(out.String())
}

func stringToStones(input string) []int {
	stoneStrings := strings.Split(strings.TrimSpace(input), " ")

	stones := []int{}
	for _, stoneString := range stoneStrings {
		stone, _ := strconv.Atoi(stoneString)
		stones = append(stones, stone)
	}

	return stones
}

func blink(stones []int) []int {
	newStones := []int{}

	for _, stone := range stones {
		stoneString := strconv.Itoa(stone)
		numDigitsInStone := len(stoneString)

		if stone == 0 {
			newStones = append(newStones, 1)
		} else if numDigitsInStone%2 == 0 {
			leftHalf := stoneString[:numDigitsInStone/2]
			rightHalf := stoneString[numDigitsInStone/2:]
			leftHalfNum, _ := strconv.Atoi(leftHalf)
			rightHalfNum, _ := strconv.Atoi(rightHalf)
			newStones = append(newStones, leftHalfNum, rightHalfNum)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}

	return newStones
}

func main() {
	input, err := readFile("../aoc24-11-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := time.Now()
	newStones := stringToStones(input)
	for range 40 {
		newStones = blink(newStones)
	}
	fmt.Println(time.Since(t))

	fmt.Println(len(newStones))
}

func readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(fileContent), nil
}
