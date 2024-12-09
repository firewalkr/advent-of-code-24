package main

import (
	"testing"
)

func Test_isReportSafe(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output bool
	}{
		{
			name:   "levels decreasing by 1 to 3 are safe",
			input:  "7 6 4 2 1",
			output: true,
		},
		{
			name:   "levels decreasing by > 3 are not safe",
			input:  "9 7 6 2 1",
			output: false,
		},
		{
			name:   "levels increasing by 1 to 3 are safe",
			input:  "1 3 6 7 9",
			output: true,
		},
		{
			name:   "levels increasing by > 3 are not safe",
			input:  "1 2 7 8 9",
			output: false,
		},
		{
			name:   "levels with same value are not safe",
			input:  "8 6 4 4 1",
			output: false,
		},
		{
			name:   "mix of increasing and decreasing levels are not safe",
			input:  "1 3 2 4 5",
			output: false,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := isReportSafe(c.input)
		if got != c.output {
			t.Errorf("ReportSafety(%q) == %t, want %t", c.name, got, c.output)
		}
	}
}
