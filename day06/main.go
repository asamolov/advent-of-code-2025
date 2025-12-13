package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

func add(agg int, item int) int {
	fmt.Printf("+%d", item)
	return agg + item
}

func mul(agg int, item int) int {
	fmt.Printf("*%d", item)
	return agg * item
}

func Slice(lines []string, idx int) string {
	var sb strings.Builder
	for _, l := range lines {
		sb.WriteByte(l[idx])
	}
	return sb.String()
}

func ColInt(lines []string, idx int) int {
	s := strings.TrimSpace(Slice(lines, idx))
	i, _ := strconv.Atoi(s)
	return i
}

func allSpaces(lines []string, idx int) bool {
	for _, l := range lines {
		if l[idx] != ' ' {
			return false
		}
	}
	return true
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")

	result := 0
	result2 := 0
	idx := 0
	args := make([][]int, len(lines)-1)
	for idx = 0; idx < len(lines)-1; idx++ {
		line := lines[idx]
		for _, arg := range strings.Fields(line) {
			p, _ := strconv.Atoi(arg)
			args[idx] = append(args[idx], p)
		}
	}

	var ops []string
	for _, op := range strings.Fields(lines[len(lines)-1]) {
		ops = append(ops, op)
	}

	// sanity check
	for idx, arg := range args {
		if len(arg) != len(ops) {
			panic(fmt.Sprintf("length of ops and arg line %d do not match!", idx))
		}
	}

	// calculate
	for idx, op := range ops {
		fn := add
		acc := 0
		if op == "*" {
			fn = mul
			acc = 1
		}
		for _, arg := range args {
			acc = fn(acc, arg[idx])
		}
		fmt.Printf("=%d\n", acc)
		result += acc
	}

	fmt.Println("Part 2")
	opIdx := 0
	acc := 0
	fn := add
	if ops[opIdx] == "*" {
		fn = mul
		acc = 1
	}
	argLines := lines[:len(lines)-1]
	// transposed problems one by one
	for i := 0; i < len(lines[0]); i++ {
		if allSpaces(lines, i) {
			fmt.Printf("=%d\n", acc)
			result2 += acc
			// start new problem - reset accumulator and select fn
			opIdx++
			op := ops[opIdx]
			fn = add
			acc = 0
			if op == "*" {
				fn = mul
				acc = 1
			}
		} else {
			acc = fn(acc, ColInt(argLines, i))
		}
	}
	// final acc
	fmt.Printf("=%d\n", acc)
	result2 += acc

	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
