package main

// this one takes some 10 secs on an M1 Max.
// will revisit.

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MissingOp struct {
	Total     uint64
	Operands  []uint64
	operators []string
}

func (m MissingOp) String() string {
	if len(m.operators) == 0 {
		s := make([]string, len(m.Operands))
		for i, op := range m.Operands {
			s[i] = strconv.FormatUint(op, 10)
		}
		return strings.Join(s, " ")
	}

	buf := strings.Builder{}
	buf.WriteString(strconv.FormatUint(m.Operands[0], 10))
	for i, op := range m.operators {
		buf.WriteString(" ")
		buf.WriteString(op)
		buf.WriteString(" ")
		buf.WriteString(strconv.FormatUint(m.Operands[i+1], 10))
	}

	buf.WriteString(" = ")
	buf.WriteString(strconv.FormatUint(m.Total, 10))

	return buf.String()
}

// if it was binary, 100 (4), 3rd digit from right
// 4 / 2 == 2 -> 4 % 2 == 0
// 2 / 2 == 1 -> 2 % 2 == 0
// 1 / 2 == 0 -> 1 % 2 == 1
//
// for ternary, e.g. 210 (21), 3rd digit from right
// 21 / 3 == 7 -> 21 % 3 == 0
//
//	7 / 3 == 2 -> 7 % 3 == 1
//	2 / 3 == 0 -> 2 % 3 == 2
//
// this assumes i always has > x ternary digits
func ternaryNthDigit(i, x int) int {
	var mod int
	for ; x > 0; x-- {
		div := i / 3
		mod = i % 3
		i = div
	}

	return mod
}

func IsTotalFeasible(missingOp *MissingOp) bool {
	numSpaces := len(missingOp.Operands) - 1
	// 0 will be x, 1 will be +, 2 will be ||
	for i := 0; i < int(math.Pow(3, float64(numSpaces))); i++ {
		// left to right, e.g. "01012" -> "11 x 22 + 33 x 44 + 55 || 66" where || concatenates the digits of both sides
		// e.g. 12 || 23 = 1223
		acc := missingOp.Operands[0]
		operators := make([]string, 0)

		for s := range numSpaces {
			v := ternaryNthDigit(i, s+1)

			if v == 0 {
				operators = append(operators, "x")
				acc *= missingOp.Operands[s+1]
			} else if v == 1 {
				operators = append(operators, "+")
				acc += missingOp.Operands[s+1]
			} else {
				operators = append(operators, "||")
				numDigitsRightSide := func(n uint64) int {
					if n == 0 {
						return 1
					}
					return int(math.Log10(float64(n))) + 1
				}(missingOp.Operands[s+1])
				acc = acc*uint64(math.Pow10(numDigitsRightSide)) + missingOp.Operands[s+1]
			}

			// if acc > missingOp.Total {
			// 	missingOp.operators = operators
			// 	break
			// }
		}
		if acc == missingOp.Total {
			missingOp.operators = operators
			// fmt.Println("acc:", acc, "op", missingOp)

			return true
		}
	}

	return false
}

func main() {
	missingOps, err := readFile("../aoc24-7-input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum := uint64(0)
	for _, missingOp := range missingOps {
		if IsTotalFeasible(&missingOp) {
			sum += missingOp.Total
		}
	}

	fmt.Println(sum)
}

func readFile(filename string) ([]MissingOp, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	missingOps := []MissingOp{}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		totalAndOperands := strings.Split(line, ":")
		total, _ := strconv.ParseUint(totalAndOperands[0], 10, 64)
		operands := strings.Split(strings.TrimSpace(totalAndOperands[1]), " ")
		ops := make([]uint64, len(operands))
		for i, op := range operands {
			ops[i], _ = strconv.ParseUint(op, 10, 64)
		}

		missingOps = append(missingOps, MissingOp{
			Total:    total,
			Operands: ops,
		})
	}

	return missingOps, nil
}
