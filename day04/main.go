package main

import (
	"fmt"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

type Field struct {
	lines         []string
	height, width int
}

func newField(lines []string) Field {
	return Field{lines: lines, height: len(lines), width: len(lines[0])}
}

func (f *Field) nRolls(x, y int) int {
	if x < 0 || x >= f.width || y < 0 || y >= f.height {
		return 0
	}
	if f.lines[y][x] == '@' {
		return 1
	}
	return 0
}

func (f *Field) accessibleRolls() int {
	result := 0
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			if f.nRolls(x, y) == 0 {
				continue
			}
			rolls := 0
			rolls += f.nRolls(x-1, y-1) + f.nRolls(x, y-1) + f.nRolls(x+1, y-1)
			rolls += f.nRolls(x-1, y) + f.nRolls(x+1, y)
			rolls += f.nRolls(x-1, y+1) + f.nRolls(x, y+1) + f.nRolls(x+1, y+1)
			if rolls < 4 {
				result++
			}
		}
	}
	return result
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")
	result := 0
	result2 := 0
	f := newField(lines)
	result = f.accessibleRolls()
	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
