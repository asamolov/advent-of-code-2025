package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

func parseInterval(interval string) (*Interval, error) {
	x := strings.Split(interval, "-")
	if len(x) != 2 {
		return nil, fmt.Errorf("can't parse interval from '%s'", interval)
	}
	start, err := strconv.Atoi(x[0])
	if err != nil {
		return nil, fmt.Errorf("can't parse interval from '%s': [%w]", interval, err)
	}
	finish, err := strconv.Atoi(x[1])
	if err != nil {
		return nil, fmt.Errorf("can't parse interval from '%s': [%w]", interval, err)
	}
	return &Interval{start: start, finish: finish}, nil
}

type Interval struct {
	start, finish int
}

func (interval *Interval) Contains(i int) bool {
	return i >= interval.start && i <= interval.finish
}

func (interval *Interval) Size() int {
	return interval.finish - interval.start + 1
}
func (interval *Interval) String() string {
	return fmt.Sprintf("%d-%d", interval.start, interval.finish)
}

func (interval *Interval) Overlap(other *Interval) *Interval {
	if interval.Contains(other.start) || interval.Contains(other.finish) ||
		other.Contains(interval.start) || other.Contains(interval.finish) {
		return &Interval{start: min(interval.start, other.start), finish: max(interval.finish, other.finish)}
	}
	return nil
}

func checkRanges(ranges []*Interval, i int) (*Interval, bool) {
	for _, r := range ranges {
		if r.Contains(i) {
			return r, true
		}
	}
	return nil, false
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")

	result := 0
	result2 := 0
	idx := 0

	var ranges []*Interval
	for idx = 0; idx < len(lines); idx++ {
		line := lines[idx]
		if len(line) == 0 {
			break
		}
		interval, err := parseInterval(line)
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, interval)
	}

	for idx++; idx < len(lines); idx++ {
		line := lines[idx]
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		_, inAnyRange := checkRanges(ranges, i)
		if inAnyRange {
			result++
		}
	}

	// merge intervals in-place
	mergeIntervals(ranges)

	// calculate sizes
	for _, i := range ranges {
		if i != nil {
			fmt.Printf("Merged interval: %s\n", i)
			result2 += i.Size()
		}
	}

	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
}

func mergeIntervals(ranges []*Interval) (hasMerged bool) {
	for idx, r := range ranges {
		if r == nil {
			continue
		}
		for i := idx + 1; i < len(ranges); i++ {
			overlapCandidate := ranges[i]
			if overlapCandidate == nil {
				continue
			}
			newInterval := r.Overlap(overlapCandidate)
			if newInterval != nil {
				fmt.Printf("Merging %s + %s => %s\n", r, overlapCandidate, newInterval)
				ranges[i] = newInterval
				ranges[idx] = nil
				hasMerged = true
				break
			}
		}
	}
	return
}
