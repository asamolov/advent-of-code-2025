package main

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

type JunctionBox struct {
	x, y, z int
}

type Cable struct {
	left, right JunctionBox
	distance    int
}

func newCable(left, rigth JunctionBox) Cable {
	return Cable{left: left, right: rigth, distance: left.distance(rigth)}
}

func parseBox(line string) JunctionBox {
	coords := strings.Split(line, ",")
	var box JunctionBox
	box.x, _ = strconv.Atoi(coords[0])
	box.y, _ = strconv.Atoi(coords[1])
	box.z, _ = strconv.Atoi(coords[2])
	return box
}

func (box *JunctionBox) distance(other JunctionBox) int {
	dx := box.x - other.x
	dy := box.y - other.y
	dz := box.z - other.z
	return dx*dx + dy*dy + dz*dz
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")

	var boxes []JunctionBox
	for _, line := range lines {
		boxes = append(boxes, parseBox(line))
	}

	var cables []Cable
	for i, b1 := range boxes {
		for j, b2 := range boxes {
			if i >= j {
				// only one way cables
				continue
			}
			cables = append(cables, newCable(b1, b2))
		}
	}

	sort.Slice(cables, func(i, j int) bool {
		return cables[i].distance < cables[j].distance
	})

	// junction -> circuit
	idx := make(map[JunctionBox]*[]JunctionBox)
	// prefill
	for _, j := range boxes {
		idx[j] = &[]JunctionBox{j}
	}

	limitCables := 10
	if len(boxes) > 100 {
		limitCables = 1000
	}
	for i, c := range cables {
		if i >= limitCables {
			break
		}
		// merge circuits
		leftC := idx[c.left]
		rightC := idx[c.right]
		// update pointers to new circuit
		fmt.Printf("Cable %v\n", c)
		if leftC != rightC {
			mergedCircuit := make([]JunctionBox, 0)
			mergedCircuit = append(mergedCircuit, *leftC...)
			mergedCircuit = append(mergedCircuit, *rightC...)
			fmt.Printf("M: %v + %v => %v\n", *leftC, *rightC, mergedCircuit)
			for _, j := range mergedCircuit {
				idx[j] = &mergedCircuit
			}
		}
	}

	// collect all circuits
	circuits := slices.Collect(maps.Values(idx))
	sort.Slice(circuits, func(i, j int) bool {
		pI := circuits[i]
		pJ := circuits[j]
		return len(*pI) > len(*pJ)
	})
	circuits = slices.Compact(circuits)
	fmt.Printf("Circuits:\n")
	for _, c := range circuits {
		fmt.Printf("%v\n", *c)
	}

	// take and multiply sizes of first three
	result := 1
	result2 := 0

	for _, c := range circuits[:3] {
		result *= len(*c)
	}

	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}
