package main

import (
	"fmt"
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

func IsTotalFeasible(missingOp *MissingOp) bool {
	numSpaces := len(missingOp.Operands) - 1
	// 0 will be x, 1 will be +
	for i := 0; i < 2<<numSpaces; i++ {
		// left to right, e.g. "0101" -> "x + x +"
		acc := missingOp.Operands[0]
		operators := make([]string, numSpaces)

		for s := range numSpaces {
			v := i & (1 << (numSpaces - s))
			// or 1 & int(math.Pow(2, float64(numSpaces-s)))

			if v == 0 {
				operators[s] = "x"
				acc *= missingOp.Operands[s+1]
			} else {
				operators[s] = "+"
				acc += missingOp.Operands[s+1]
			}

			if acc > missingOp.Total {
				missingOp.operators = operators
				break
			}
		}

		if acc == missingOp.Total {
			missingOp.operators = operators
			fmt.Println(missingOp)
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
