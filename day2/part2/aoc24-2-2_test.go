package main

import (
	"testing"
)

func Test_isReportSafeWithAtMostOneFlaw(t *testing.T) {
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
			name:   "same level only once is safe",
			input:  "8 6 4 4 1",
			output: true,
		},
		{
			name:   "inversion only happening once at the beginning is safe",
			input:  "3 1 2 4 5",
			output: true,
		},
		{
			name:   "inversion only happening once in the middle is safe",
			input:  "8 6 7 4 2",
			output: true,
		},
		{
			name:   "inversion only happening once at the end is safe",
			input:  "1 3 5 8 6",
			output: true,
		},
		{
			name:   "inversion happening more than once is not safe",
			input:  "1 3 2 5 4",
			output: false,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := isReportSafeWithAtMostOneFlaw_Ugly(c.input)
		if got != c.output {
			t.Errorf("ReportSafety(%q) == %t, want %t", c.name, got, c.output)
		}
	}
}
