package main

import (
	"fmt"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

type Field struct {
	lines         [][]byte
	timelines     [][]int
	height, width int
}

func newField(strings []string) Field {
	bytes := make([][]byte, 0, len(strings))
	timelines := make([][]int, 0, len(strings))
	for _, l := range strings {
		bytes = append(bytes, []byte(l))
		timelines = append(timelines, make([]int, len(l)))
	}
	return Field{lines: bytes, timelines: timelines, height: len(bytes), width: len(bytes[0])}
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

func (f *Field) has(x, y int, ch byte) bool {
	if x < 0 || x >= f.width || y < 0 || y >= f.height {
		return false
	}
	return f.lines[y][x] == ch
}

func (f *Field) findStart() int {
	for i, ch := range f.lines[0] {
		if ch == 'S' {
			return i
		}
	}
	return -1
}

func (f *Field) propagateBeam(level int) (int, *Field) {
	next := f.clone()
	splits := 0
	for i, ch := range f.lines[level] {
		if ch == '|' || ch == 'S' {
			// check splitter
			if next.lines[level+1][i] == '^' {
				next.lines[level+1][i-1] = '|'
				next.lines[level+1][i+1] = '|'
				splits++
			} else {
				next.lines[level+1][i] = '|'
			}
		}
	}
	return splits, &next
}

func (f *Field) findTimelinesDfs(level, pos int) int {
	level++
	if level >= f.height {
		return 1
	}

	timelines := 0
	ch := f.lines[level][pos]
	if ch == '^' {
		timelines += f.findTimelinesDfs(level, pos-1)
		timelines += f.findTimelinesDfs(level, pos+1)
	} else {
		timelines += f.findTimelinesDfs(level, pos)
	}
	return timelines
}

func (f *Field) findTimelinesBfs() int {
	timelinesAtLevel := 0
	f.timelines[0][f.findStart()] = 1
	for level := 1; level < f.height; level++ {
		timelinesAtLevel = 0
		for i, ch := range f.lines[level] {
			timelines := f.timelines[level-1][i]
			if ch == '^' {
				f.timelines[level][i-1] += timelines
				f.timelines[level][i+1] += timelines
				timelinesAtLevel += 2 * timelines
			} else {
				f.timelines[level][i] += timelines
				timelinesAtLevel += timelines
			}
		}
		fmt.Printf("%4d: %d\n", level, timelinesAtLevel)
	}
	return timelinesAtLevel
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")
	result := 0
	result2 := 0
	f := newField(lines)
	for level := 0; level < f.height-1; level++ {
		splits, newField := f.propagateBeam(level)
		result += splits
		f = *newField
	}
	f = newField(lines)
	result2 += f.findTimelinesBfs()
	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
