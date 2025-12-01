package main

import (
	"fmt"
	"strconv"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")
	pos := 50
	result := 0
	result2 := 0
	for i, l := range lines {
		fmt.Printf("cmd %3d: %s\n", i, l)
		dir := l[0]
		n, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(fmt.Sprintf("cannot parse line %d: %s", i, l))
		}
		if dir == 'L' {
			n = -n
		}
		oldPos := pos
		pos += n
		result2 += utils.Abs(pos / 100)
		pos %= 100
		if pos < 0 {
			pos += 100
			if oldPos != 0 {
				result2++
			}
		}
		if pos == 0 {
			result += 1
			if n < 0 {
				result2 += 1
			}
		}
		fmt.Printf("pos: %3d, dir: %3d, res1: %d, res2: %d\n", pos, n, result, result2)
	}
	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
