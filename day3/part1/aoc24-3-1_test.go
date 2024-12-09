package main

import (
	"testing"
)

func Test_addMuls(t *testing.T) {
	// Test cases
	cases := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "adds muls correctly",
			input:  "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			output: 2*4 + 5*5 + 11*8 + 8*5, // 161
		},
		{
			name:   "adds bigger muls correctly",
			input:  "[#from())when()/}+%mul(982,733)mul(700,428)}}dont(){:,$+mul(395,45)[",
			output: 982*733 + 700*428 + 395*45, // 719,306 + 299,600 + 17,775 = 1,036,681
		},
		{
			name:   "respects the last don't()",
			input:  "[#from())when()/}+%mul(982,733)mul(700,428)}}don't(){:,$+mul(395,45)[",
			output: 982*733 + 700*428, // 719,306 + 299,600 = 1,018,906,
		},
		{
			name:   "respects do() and don't()",
			input:  "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			output: 2*4 + 8*5, // 8 + 40 = 48
		},
	}

	// Execute test cases
	for _, c := range cases {
		got := addMuls(c.input)
		if got != c.output {
			t.Errorf("ReportSafety(%q) == %d, want %d", c.name, got, c.output)
		}
	}
}
