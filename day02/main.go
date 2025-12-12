package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

func parseInterval(interval string) (int, int, error) {
	x := strings.Split(interval, "-")
	if len(x) != 2 {
		return 0, 0, fmt.Errorf("can't parse interval from '%s'", interval)
	}
	start, err := strconv.Atoi(x[0])
	if err != nil {
		return 0, 0, fmt.Errorf("can't parse interval from '%s': [%w]", interval, err)
	}
	finish, err := strconv.Atoi(x[1])
	if err != nil {
		return 0, 0, fmt.Errorf("can't parse interval from '%s': [%w]", interval, err)
	}
	return start, finish, nil
}

func isDupe(x int) bool {
	s := strconv.Itoa(x)
	left := s[:len(s)/2]
	right := s[len(s)/2:]
	return left == right
}

func isRepeated(x int) bool {
	s := ([]byte)(strconv.Itoa(x))
out:
	for i := 1; i < len(s)/2+1; i++ {
		if len(s)%i != 0 {
			// short-cut if i is not divider of len(s)
			continue
		}
		var first []byte
		for c := range slices.Chunk(s, i) {
			if first == nil {
				first = c
			} else {
				if !slices.Equal(first, c) {
					continue out
				}
			}
		}
		return true
	}
	return false
}

func findIn(start, end int, predicate func(int) bool) (result []int) {
	result = []int{}

	for x := start; x <= end; x++ {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")
	line := lines[0]

	result := 0
	result2 := 0
	for _, interval := range strings.Split(line, ",") {
		fmt.Printf("%s\n", interval)
		start, end, err := parseInterval(interval)
		if err != nil {
			panic(err)
		}
		for _, d := range findIn(start, end, isDupe) {
			fmt.Printf("\tD: %d\n", d)
			result += d
		}
		for _, d := range findIn(start, end, isRepeated) {
			fmt.Printf("\tR: %d\n", d)
			result2 += d
		}

	}
	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
