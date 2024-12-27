package main

import (
	"strings"
	"testing"
)

var aocSample1 = strings.TrimSpace(`
Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=10000000012748, Y=10000000012176
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
		output  int64
	}{
		{
			name:    "aoc sample 1",
			machine: parseTestMachine(aocSample1),
			output:  459236326669,
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

var mat1 = Mat{
	NewVec(3, 4, 5),
	NewVec(1, 2, 1),
}

var mat2 = Mat{
	NewVec(94, 22, 8400),
	NewVec(34, 67, 5400),
}

func Test_solve(t *testing.T) {
	cases := []struct {
		name   string
		mat    Mat
		output Vec
	}{
		{
			name:   "simple system with 2 linear eqs",
			mat:    mat1,
			output: NewVec(3, -1),
		},
		{
			name:   "aoc sample 1 in linear eq form",
			mat:    mat2,
			output: NewVec(80, 40),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := solve(c.mat)
			if !out.equal(c.output) {
				t.Errorf("Expected %s, but got %s", c.output, out)
			}
		})
	}
}
