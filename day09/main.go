package main

import (
	"cmp"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/asamolov/advent-of-code-2025/internal/utils"
)

type Field struct {
	lines         [][]byte
	height, width int
	img           *image.RGBA
}

type Tile struct {
	x, y int
}

type Rect struct {
	left, right Tile
	area        int
}

func newRect(left, rigth Tile) Rect {
	return Rect{left: left, right: rigth, area: left.area(rigth)}
}

func parseTile(line string) Tile {
	coords := strings.Split(line, ",")
	var box Tile
	box.x, _ = strconv.Atoi(coords[0])
	box.y, _ = strconv.Atoi(coords[1])
	return box
}

func (box *Tile) area(other Tile) int {
	dx := utils.Abs(box.x-other.x) + 1
	dy := utils.Abs(box.y-other.y) + 1
	return dx * dy
}

func makeField(tiles []Tile) Field {
	width := slices.MaxFunc(tiles, func(a, b Tile) int {
		return cmp.Compare(a.x, b.x)
	}).x + 2

	height := slices.MaxFunc(tiles, func(a, b Tile) int {
		return cmp.Compare(a.y, b.y)
	}).y + 2
	lines := make([][]byte, height)
	for line := range lines {
		lines[line] = slices.Repeat([]byte{'.'}, width)
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	return Field{lines: lines, img: img, width: width, height: height}
}

func (f *Field) has(pt Tile, ch byte) bool {
	if pt.x < 0 || pt.x >= f.width || pt.y < 0 || pt.y >= f.height {
		return false
	}
	return f.lines[pt.y][pt.x] == ch
}

func (f *Field) set(pt Tile, ch byte) {
	if pt.x < 0 || pt.x >= f.width || pt.y < 0 || pt.y >= f.height {
		return
	}
	f.lines[pt.y][pt.x] = ch
	cl := color.RGBA{0, 0, 0, 0xff}
	switch ch {
	case 'R':
		cl = color.RGBA{128, 0, 0, 0xFF}
	case 'G':
		cl = color.RGBA{0, 128, 0, 0xFF}
	case 'X':
		cl = color.RGBA{0, 0, 128, 0xFF}
	}
	f.img.Set(pt.x, pt.y, cl)
}

func (f *Field) fillLine(start, end Tile) {
	dx := cmp.Compare(end.x, start.x)
	dy := cmp.Compare(end.y, start.y)
	f.set(start, 'R')
	f.set(end, 'R')
	curr := start
	for {
		curr.x += dx
		curr.y += dy
		if curr == end {
			break
		}
		f.set(curr, 'G')
	}
}

func (f *Field) String() string {
	if f.height > 100 || f.width > 100 {
		return fmt.Sprintf("field[%dx%d] is too big!", f.width, f.height)
	}
	var sb strings.Builder
	for _, line := range f.lines {
		sb.Write(line)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (f *Field) floodFill(start Tile, from, to byte) {
	q := make([]Tile, 0)
	q = append(q, start)

	for len(q) > 0 {
		// poll
		pt := q[0]
		q = q[1:]

		if f.has(pt, from) {
			f.set(pt, to)
			// adding left/rigth/top/bottom
			q = append(q,
				Tile{pt.x, pt.y + 1},
				Tile{pt.x, pt.y - 1},
				Tile{pt.x + 1, pt.y},
				Tile{pt.x - 1, pt.y})
		}
	}
}

func (f *Field) checkLine(start, end Tile) bool {
	dx := cmp.Compare(end.x, start.x)
	dy := cmp.Compare(end.y, start.y)

	if f.has(start, 'X') {
		return false
	}
	curr := start
	for {
		curr.x += dx
		curr.y += dy
		if f.has(curr, 'X') {
			return false
		}
		if curr == end {
			break
		}
	}
	return true
}

func (f *Field) validRect(start, end Tile) bool {
	return f.checkLine(start, Tile{start.x, end.y}) &&
		f.checkLine(start, Tile{end.x, start.y}) &&
		f.checkLine(end, Tile{start.x, end.y}) &&
		f.checkLine(end, Tile{end.x, start.y})
}

func (f *Field) drawRect(start, end Tile) {
	f.fillLine(start, Tile{start.x, end.y})
	f.fillLine(start, Tile{end.x, start.y})
	f.fillLine(end, Tile{start.x, end.y})
	f.fillLine(end, Tile{end.x, start.y})
}

func compressTiles(tiles []Tile) (compressed []Tile, mapX map[int]int, mapY map[int]int) {
	temp := make([]Tile, 0)
	temp = append(temp, tiles...)
	mapX = make(map[int]int)
	mapY = make(map[int]int)
	// sort by X
	slices.SortFunc(temp, func(left, right Tile) int {
		return cmp.Compare(left.x, right.x)
	})
	// map X
	x := 0
	for idx, t := range temp {
		mappedX, ok := mapX[t.x]
		if !ok {
			_, adjanced := mapX[t.x-1]
			if adjanced {
				x += 1
			} else {
				x += 2 // to  avoid joining
			}
			mappedX = x
			mapX[t.x] = x
		}
		t.x = mappedX
		temp[idx] = t
	}
	// sort by Y
	slices.SortFunc(temp, func(left, right Tile) int {
		return cmp.Compare(left.y, right.y)
	})
	// map Y
	y := 0
	for idx, t := range temp {
		mappedY, ok := mapY[t.y]
		if !ok {
			_, adjanced := mapY[t.y-1]
			if adjanced {
				y += 1
			} else {
				y += 2 // to  avoid joining
			}
			mappedY = y
			mapY[t.y] = y
		}
		t.y = mappedY
		temp[idx] = t
	}
	compressed = append(compressed, tiles...)
	for i := 0; i < len(compressed); i++ {
		compressed[i] = compressed[i].compress(mapX, mapY)
	}
	return
}

func (t *Tile) compress(mapX, mapY map[int]int) Tile {
	return Tile{x: mapX[t.x], y: mapY[t.y]}
}

func main() {
	defer utils.Timer("task")()

	lines := utils.ReadInput()
	fmt.Println("reading input")

	var tiles []Tile
	for _, line := range lines {
		tiles = append(tiles, parseTile(line))
	}

	var rects []Rect
	for i, b1 := range tiles {
		for j, b2 := range tiles {
			if i >= j {
				// only one way rects
				continue
			}
			rects = append(rects, newRect(b1, b2))
		}
	}

	sort.Slice(rects, func(i, j int) bool {
		return rects[i].area > rects[j].area
	})

	result := rects[0].area

	// part 2
	// compress
	compressed, mapX, mapY := compressTiles(tiles)
	field := makeField(compressed)
	start := compressed[0]
	for i := 1; i < len(compressed); i++ {
		end := compressed[i]
		field.fillLine(start, end)
		start = end
	}
	field.fillLine(start, compressed[0])
	fmt.Println("... filled lines")
	fmt.Printf("All Lines:\n%s", &field)
	field.floodFill(Tile{0, 0}, '.', 'X')
	fmt.Println("... flood filled")
	fmt.Printf("After Fill:\n%s", &field)
	result2 := 0
	for idx, rect := range rects {
		if field.validRect(rect.left.compress(mapX, mapY),
			rect.right.compress(mapX, mapY)) {
			result2 = rect.area
			fmt.Printf("%d-th rect is valid!\n", idx)
			field.drawRect(rect.left.compress(mapX, mapY),
				rect.right.compress(mapX, mapY))
			break
		}
	}

	fmt.Printf("result: %d\n", result)
	fmt.Printf("result2: %d\n", result2)
	f, _ := os.Create("image.png")
	png.Encode(f, field.img)
}
