package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`)

func Test_sumNumberOfMiddlePageOfCorrectUpdates(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "aoc sample",
			input:  testInput,
			output: 143,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := sumNumberOfMiddlePageOfCorrectUpdates(c.input)
		if got != c.output {
			t.Errorf("ReportSafety(%q) == %d, want %d", c.name, got, c.output)
		}
	}
}
