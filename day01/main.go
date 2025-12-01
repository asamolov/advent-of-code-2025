package main

import (
	"fmt"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

func main() {
	defer utils.Timer("task")()
	lines := utils.ReadInput()
	fmt.Println("reading input")
	for i, l := range lines {
		fmt.Printf("%3d: %s\n", i, l)
	}
}
