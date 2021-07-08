package main

// Describe Shape for 2D and 3D shapes

type Shape interface {
	Area() float64
	Scale(float64)
}

type Shape3D interface {
	Shape
	Volume() float64
}

type Cube struct {
	Side float64
}

func (c Cube) Area() float64 {
	return 6 * c.Side * c.Side
}

func (c *Cube) Scale(scaleBy float64) {
	c.Side *= scaleBy
}

func (c Cube) Volume() float64 {
	return c.Side * c.Side * c.Side
}

func main() {

}
