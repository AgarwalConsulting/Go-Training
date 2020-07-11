package main

// Cube conforms to Shape and Shape3D
type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * c.side * c.side // Surface area of cube
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}
