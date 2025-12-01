package main

import (
	"fmt"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

func main() {
	lines := utils.readInput()
	fmt.Println("reading input")
	for i, l := range lines {
		fmt.Println("%3d: %s", i, l)
	}
}
