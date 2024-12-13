package main

import (
	"testing"
)

func Test_IsTotalFeasible(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  MissingOp
		output bool
	}{
		{
			name: "sample 0",
			input: MissingOp{
				Total:    190,
				Operands: []uint64{1, 90},
			},
			output: true,
		},
		{
			name: "sample 1",
			input: MissingOp{
				Total:    190,
				Operands: []uint64{10, 19},
			},
			output: true,
		},
		{
			name: "sample 2",
			input: MissingOp{
				Total:    3267,
				Operands: []uint64{81, 40, 27},
			},
			output: true,
		},
		{
			name: "sample 3",
			input: MissingOp{
				Total:    83,
				Operands: []uint64{17, 5},
			},
			output: false,
		},
		{
			name: "sample 4",
			input: MissingOp{
				Total:    156,
				Operands: []uint64{15, 6},
			},
			output: true,
		},
		{
			name: "sample 5",
			input: MissingOp{
				Total:    7290,
				Operands: []uint64{6, 8, 6, 15},
			},
			output: true,
		},
		{
			name: "sample 6",
			input: MissingOp{
				Total:    161011,
				Operands: []uint64{16, 10, 13},
			},
			output: false,
		},
		{
			name: "sample 7",
			input: MissingOp{
				Total:    192,
				Operands: []uint64{17, 8, 14},
			},
			output: true,
		},
		{
			name: "sample 8",
			input: MissingOp{
				Total:    21037,
				Operands: []uint64{9, 7, 18, 13},
			},
			output: false,
		},
		{
			name: "sample 9",
			input: MissingOp{
				Total:    292,
				Operands: []uint64{11, 6, 16, 20},
			},
			output: true,
		},
		{
			name: "works on base 2",
			input: MissingOp{
				Total:    75669678,
				Operands: []uint64{7, 5, 7, 41, 1, 53, 6, 239, 414},
			},
			output: true,
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := IsTotalFeasible(&c.input)
		if got != c.output {
			t.Errorf("IsTotalFeasible(%q) == %v, want %v", c.name, got, c.output)
		}
	}
}
