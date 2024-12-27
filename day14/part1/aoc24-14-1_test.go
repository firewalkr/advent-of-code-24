package main

import (
	"strings"
	"testing"
)

var aocSample1 = strings.TrimSpace(`
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`)

func Test_calcSafetyFactorAfter(t *testing.T) {
	type args struct {
		bots      []*Bot
		numSteps  int
		gridSizeX int
		gridSizeY int
	}

	cases := []struct {
		name   string
		args   args
		output int
	}{
		{
			name: "aoc sample 1",
			args: args{
				bots:      readBots(aocSample1),
				numSteps:  100,
				gridSizeX: 11,
				gridSizeY: 7,
			},
			output: 12,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := calcSafetyFactorAfter(c.args.bots, c.args.numSteps, c.args.gridSizeX, c.args.gridSizeY)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}
