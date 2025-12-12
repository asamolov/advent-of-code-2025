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

func goDeep(m map[int][]int, maxIdx, curr, startIdx, rem int) (result int, found bool) {
	rem--
	if rem < 0 {
		return curr, true
	}
	for j := 9; j >= 1; j-- {
		js := m[j]
		idx := slices.IndexFunc(js, func(jidx int) bool {
			return jidx >= startIdx && jidx < maxIdx-rem
		})
		if idx >= 0 {
			result, found := goDeep(m, maxIdx, curr*10+j, js[idx]+1, rem)
			if found {
				return result, found
			}
		}
	}
	return 0, false
}

func findMaxJolt12(m map[int][]int, maxIdx int) int {
	result, _ := goDeep(m, maxIdx, 0, 0, 12)
	return result
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
		maxJolt2, _ := goDeep(m, len(l), 0, 0, 2)
		maxJolt12 := findMaxJolt12(m, len(l))
		fmt.Printf("%s: %d\n", l, maxJolt)
		fmt.Printf("%s: %d\n", l, maxJolt2)
		fmt.Printf("%s: %d\n", l, maxJolt12)
		result += maxJolt
		result2 += maxJolt12
	}
	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
