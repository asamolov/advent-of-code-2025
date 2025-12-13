package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

func add(agg int, item int) int {
	return agg + item
}

func mul(agg int, item int) int {
	return agg * item
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
		result += acc
	}

	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
