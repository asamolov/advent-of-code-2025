package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

type Task struct {
	target  int
	buttons []int
}

func parseTask(line string) Task {
	var t Task
	for _, group := range strings.Fields(line) {
		switch group[0] {
		case '[':
			t.target = 0
			mask := 1
			for _, ch := range group[1 : len(group)-1] {
				if ch == '#' {
					t.target |= mask
				}
				mask <<= 1
			}
		case '(':
			button := 0
			for _, x := range strings.Split(group[1:len(group)-1], ",") {
				num, _ := strconv.Atoi(x)
				button |= 1 << num
			}
			t.buttons = append(t.buttons, button)
		}
	}
	return t
}

func (t *Task) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "[%0b] ", t.target)
	for idx, btn := range t.buttons {
		if idx > 0 {
			fmt.Fprint(&sb, ", ")
		}
		fmt.Fprintf(&sb, "%0b", btn)
	}
	return sb.String()
}

func (t *Task) solve() int {
	for deep := 1; deep <= len(t.buttons); deep++ {
		for _, combination := range selectN(deep, t.buttons) {
			out := t.target
			for _, btn := range combination {
				out ^= btn
			}
			if out == 0 {
				return deep
			}
		}
	}
	return 0
}

func selectN(deep int, arr []int) (combinations [][]int) {
	combinations = make([][]int, 0)
	if deep == 0 {
		combinations = append(combinations, []int{})
		return
	}
	for idx, n := range arr {
		for _, selected := range selectN(deep-1, arr[idx+1:]) {
			comb := append(selected, n)
			combinations = append(combinations, comb)
		}
	}
	return
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")

	result := 0
	result2 := 0

	for idx, l := range lines {
		t := parseTask(l)
		presses := t.solve()
		fmt.Printf("T%4d in %4d: %s\n", idx, presses, &t)
		result += presses
	}

	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
