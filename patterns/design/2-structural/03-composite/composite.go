package main

import (
	"fmt"
	"strings"
)

// Composite: GraphicObject with Name, Color, Position, Move method, Children
// Composite: Neuron with ConnectTo method & NeuronLayer; and Neuroner with Iter method

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

func main() {
	face1 := &GraphicObject{"Face", "Pink", []GraphicObject{
		{"Left Eye", "Brown", []GraphicObject{*NewCircle("Left Pupil")}},
		{"Right Eye", "Brown", []GraphicObject{*NewCircle("Right Pupil")}},
	}}

	fmt.Println(face1)

	face2 := &GraphicObject{"Face", "Pink", []GraphicObject{
		{"Left Eye", "Blue", []GraphicObject{*NewCircle("Left Pupil")}},
		{"Right Eye", "Blue", []GraphicObject{*NewCircle("Right Pupil")}},
	}}

	fmt.Println(face2)

	drawing := &GraphicObject{"Background", "Black", []GraphicObject{*face1, *face2}}

	fmt.Println(drawing)
}
