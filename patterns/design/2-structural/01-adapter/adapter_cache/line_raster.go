package adapter_cache

import (
	"fmt"
	"strings"
)

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height}}}
}

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

type VectorToRasterAdapter struct {
	Points []Point
}

func NewVectorToRasterAdapter(v *VectorImage) *VectorToRasterAdapter {
	v2r := VectorToRasterAdapter{}

	for _, line := range v.Lines {
		v2r.AddLine(line)
	}

	return &v2r
}

func (v2r *VectorToRasterAdapter) AddVectorImage(v *VectorImage) {
	for _, line := range v.Lines {
		v2r.AddLine(line)
	}
}

func (v *VectorToRasterAdapter) GetPoints() []Point {
	return v.Points
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func (a *VectorToRasterAdapter) AddLine(line Line) {
	// Cache the lines by computing hash of the line
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.Points = append(a.Points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.Points = append(a.Points, Point{x, top})
		}
	}

	fmt.Println("generated", len(a.Points), "points")
}
