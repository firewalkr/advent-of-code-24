package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func stringToStones(input string) []int64 {
	stoneStrings := strings.Split(strings.TrimSpace(input), " ")

	stones := []int64{}
	for _, stoneString := range stoneStrings {
		stone, _ := strconv.ParseInt(stoneString, 10, 64)
		stones = append(stones, stone)
	}

	return stones
}

type StoneGroup struct {
	stone int64
	count int64
}

type StoneGroups []StoneGroup

func stonesToGroups(stones []int64) StoneGroups {
	var groups StoneGroups
	m := make(map[int64]int64)

	for _, stone := range stones {
		m[stone]++
	}

	for stone, count := range m {
		groups = append(groups, StoneGroup{stone, count})
	}

	return groups
}

func dedupGroups(groups StoneGroups) StoneGroups {
	var deduped StoneGroups
	m := make(map[int64]int64)

	for _, sg := range groups {
		m[sg.stone] += sg.count
	}

	for stone, count := range m {
		deduped = append(deduped, StoneGroup{stone, count})
	}

	return deduped
}

func blink(stoneGroups StoneGroups) StoneGroups {
	var newGroups StoneGroups

	for _, sg := range stoneGroups {
		if sg.stone == 0 {
			newGroups = append(newGroups, StoneGroup{1, sg.count})
		} else {
			numDigitsInStone := numDigitsInStone(sg.stone)
			if numDigitsInStone%2 == 0 {
				leftHalfNum := sg.stone / int64(math.Pow10(numDigitsInStone/2))
				rightHalfNum := sg.stone % int64(math.Pow10(numDigitsInStone/2))
				newGroups = append(newGroups,
					StoneGroup{leftHalfNum, sg.count},
					StoneGroup{rightHalfNum, sg.count},
				)
			} else {
				newGroups = append(newGroups, StoneGroup{sg.stone * 2024, sg.count})
			}
		}
	}

	return dedupGroups(newGroups)
}

func numDigitsInStone(stone int64) int {
	if stone == 0 {
		return 1
	}
	return int(math.Log10(float64(stone))) + 1
}

func main() {
	input, err := readFile("../aoc24-11-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	newStones := stringToStones(input)
	newStoneGroups := stonesToGroups(newStones)
	t := time.Now()
	totalCount := int64(0)

	for range 75 {
		newStoneGroups = blink(newStoneGroups)
	}

	for _, sg := range newStoneGroups {
		totalCount += sg.count
	}

	fmt.Println(time.Since(t))

	// fmt.Println(stonesToString(newStones))
	fmt.Println(totalCount)
}

func readFile(filename string) (string, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	return string(fileContent), nil
}
