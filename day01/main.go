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
		pos += n
		pos %= 100
		if pos < 0 {
			pos += 100
		}
		fmt.Printf("dir: %3d, pos: %3d\n", n, pos)
		if pos == 0 {
			result += 1
		}
	}
	fmt.Printf("result: %d\n", result)
}
