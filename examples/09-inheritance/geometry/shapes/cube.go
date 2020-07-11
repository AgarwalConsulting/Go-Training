package shapes

// Cube implements Shape & Shape3D
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
