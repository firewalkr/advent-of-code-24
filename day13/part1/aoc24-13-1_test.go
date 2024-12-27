package main

import (
	"strings"
	"testing"
)

var aocSample1 = strings.TrimSpace(`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400
`)

func parseTestMachine(s string) *Machine {
	m := &Machine{}
	m.buttonAX, m.buttonAY = parseButtonA(s)
	m.buttonBX, m.buttonBY = parseButtonB(s)
	m.prizeX, m.prizeY = parsePrize(s)
	return m
}

func Test_calcCheapestPrice(t *testing.T) {

	cases := []struct {
		name    string
		machine *Machine
		output  int
	}{
		{
			name:    "aoc sample 1",
			machine: parseTestMachine(aocSample1),
			output:  280,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := calcCheapestPrice(c.machine)
			if out != c.output {
				t.Errorf("Expected %v, but got %v", c.output, out)
			}
		})
	}
}
