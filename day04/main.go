package main

import (
	"fmt"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

type Field struct {
	lines         [][]byte
	height, width int
}

func newField(strings []string) Field {
	bytes := make([][]byte, 0, len(strings))
	for _, l := range strings {
		bytes = append(bytes, []byte(l))
	}
	return Field{lines: bytes, height: len(bytes), width: len(bytes[0])}
}

func (f *Field) clone() Field {
	bytes := make([][]byte, 0, len(f.lines))
	for _, l := range f.lines {
		cl := make([]byte, len(l))
		copy(cl, l)
		bytes = append(bytes, cl)
	}
	return Field{lines: bytes, height: f.height, width: f.width}
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

func (f *Field) removeRolls() (int, *Field) {
	removed := 0
	newField := f.clone()
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
				removed++
				newField.lines[y][x] = 'x'
			}
		}
	}
	return removed, &newField
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")
	result := 0
	result2 := 0
	f := newField(lines)
	result, _ = f.removeRolls()

	for {
		nRemoved, newField := f.removeRolls()
		if nRemoved == 0 {
			break
		}
		result2 += nRemoved
		f = *newField
	}
	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
