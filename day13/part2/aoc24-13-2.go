package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// naive approach is no good anymore (though it worked pretty well until
// two orders of magnitude below what step2 requires).
// let's solve the linear equations instead.

type Vec []*big.Rat

func NewVec(is ...int64) Vec {
	v := make(Vec, 0, len(is))
	for _, i := range is {
		v = append(v, new(big.Rat).SetInt64(i))
	}
	return v
}

type Mat []Vec

// solve a system of 2 linear eqs with 2 vars in matrix form:
//
//	divide first row vector by value in its first column
//	  i.e. coefficient of x on first row becomes 1
//	divide second row vector by its first value and subtract values of first vector
//	  i.e. coefficient of x on second row becomes 0
//	divide second row vector by its second value
//	  i.e. coefficient of y on second row becomes 1
//
// at this point we'll have:
//
//	x + Ay = B
//	Cy = D
//
// which becomes:
//
//	x = B - A(D/C)
//	y = D/C
//
// important: we know no coeffs are 0 at the beginning
func solve(m Mat) Vec {
	// divide first vec by first val
	v1 := m[0]
	firstV1 := v1[0]
	v1.divideBy(firstV1)

	// divide second vec by first val and subtract first vec
	v2 := m[1]
	firstV2 := v2[0]
	v2.divideBy(firstV2)
	v2.subtract(v1)

	// divide second vec by second val
	secondV2 := v2[1]
	v2.divideBy(secondV2)

	// y = v2[2] / v2[1]
	// x = v1[2] - (v1[1] * y)
	y := new(big.Rat).Quo(v2[2], v2[1])
	x := new(big.Rat).Sub(v1[2], new(big.Rat).Mul(v1[1], y))

	return Vec{x, y}
}

func (v Vec) divideBy(d *big.Rat) {
	for i := range v {
		v[i] = new(big.Rat).Quo(v[i], d)
	}
}

func (v Vec) subtract(v2 Vec) {
	for i := range v {
		v[i] = new(big.Rat).Sub(v[i], v2[i])
	}
}

// assumes same length
func (v Vec) equal(v2 Vec) bool {
	for i := range v {
		if v[i].Cmp(v2[i]) != 0 {
			return false
		}
	}

	return true
}

func (v Vec) String() string {
	b := strings.Builder{}
	b.WriteRune('[')

	vs := []string{}
	for _, vv := range v {
		vs = append(vs, vv.String())
	}

	b.WriteString(strings.Join(vs, " "))

	b.WriteRune(']')

	return b.String()
}

func (m Mat) String() string {
	ms := []string{}
	for _, mv := range m {
		ms = append(ms, mv.String())
	}

	return strings.Join(ms, "\n")
}

func machineToMat(m *Machine) Mat {
	mat := Mat{}
	mat = append(mat, Vec{
		new(big.Rat).SetInt64(m.buttonAX),
		new(big.Rat).SetInt64(m.buttonBX),
		new(big.Rat).SetInt64(m.prizeX),
	})
	mat = append(mat, Vec{
		new(big.Rat).SetInt64(m.buttonAY),
		new(big.Rat).SetInt64(m.buttonBY),
		new(big.Rat).SetInt64(m.prizeY),
	})
	return mat
}

const costA = 3
const costB = 1

var rgxButtonA = regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
var rgxButtonB = regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
var rgxPrize = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

type Machine struct {
	buttonAX, buttonAY int64
	buttonBX, buttonBY int64
	prizeX, prizeY     int64
}

var zero = new(big.Rat).SetInt64(0)

func calcCheapestPrice(m *Machine) int64 {
	vec := solve(machineToMat(m))

	for _, v := range vec {
		if !v.IsInt() || v.Cmp(zero) == -1 {
			return int64(math.MaxInt64)
		}
	}

	xNum := vec[0].Num()
	yNum := vec[1].Num()

	aPresses := new(big.Int).Abs(xNum).Int64()
	bPresses := new(big.Int).Abs(yNum).Int64()

	return aPresses*costA + bPresses*costB
}

func main() {
	machines, err := readFile("../aoc24-13-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	totalCost := int64(0)

	for _, m := range machines {
		minimum := calcCheapestPrice(m)
		if minimum != math.MaxInt {
			totalCost += minimum
		}
	}

	fmt.Println(totalCost)
}

func readFile(filename string) ([]*Machine, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	machines := []*Machine{}
	nextMachine := &Machine{}

	for _, line := range strings.Split(string(fileContent), "\n") {
		if strings.Contains(line, "Button A") {
			buttonAX, buttonAY := parseButtonA(line)
			nextMachine.buttonAX = buttonAX
			nextMachine.buttonAY = buttonAY
		} else if strings.Contains(line, "Button B") {
			buttonBX, buttonBY := parseButtonB(line)
			nextMachine.buttonBX = buttonBX
			nextMachine.buttonBY = buttonBY
		} else if strings.Contains(line, "Prize") {
			prizeX, prizeY := parsePrize(line)
			nextMachine.prizeX = 10000000000000 + prizeX
			nextMachine.prizeY = 10000000000000 + prizeY
			machines = append(machines, nextMachine)
			nextMachine = &Machine{}
		}
	}

	return machines, nil
}

func parseButtonA(line string) (int64, int64) {
	matches := rgxButtonA.FindStringSubmatch(line)
	x, _ := strconv.ParseInt(matches[1], 10, 64)
	y, _ := strconv.ParseInt(matches[2], 10, 64)
	return x, y
}

func parseButtonB(line string) (int64, int64) {
	matches := rgxButtonB.FindStringSubmatch(line)
	x, _ := strconv.ParseInt(matches[1], 10, 64)
	y, _ := strconv.ParseInt(matches[2], 10, 64)
	return x, y
}

func parsePrize(line string) (int64, int64) {
	matches := rgxPrize.FindStringSubmatch(line)
	x, _ := strconv.ParseInt(matches[1], 10, 64)
	y, _ := strconv.ParseInt(matches[2], 10, 64)
	return x, y
}
