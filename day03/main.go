package main

import (
	"fmt"
	"slices"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

func toMap(s string) map[int][]int {
	result := make(map[int][]int)
	for i, ch := range s {
		k := (int)(ch - '0')
		result[k] = append(result[k], i)
	}
	return result
}

func findMaxJolt(m map[int][]int) int {
	for i := 9; i >= 1; i-- {
		is := m[i]
		if is == nil {
			continue
		}
		iidx := is[0]
		for j := 9; j >= 1; j-- {
			js := m[j]
			if slices.ContainsFunc(js, func(jidx int) bool {
				return jidx > iidx
			}) {
				return i*10 + j
			}
		}
	}
	return 0 // impossible case
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")

	result := 0
	result2 := 0
	for _, l := range lines {
		m := toMap(l)
		maxJolt := findMaxJolt(m)
		fmt.Printf("%s: %d\n", l, maxJolt)
		result += maxJolt
	}
	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
