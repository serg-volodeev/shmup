package shape

type Circle struct {
	x, y   float64
	radius float64
}

func NewCircle(x, y, radius float64) *Circle {
	return &Circle{x: x, y: y, radius: radius}
}

func (c *Circle) SetCenter(x, y float64) {
	c.x = x
	c.y = y
}

func square(n float64) float64 {
	return n * n
}

func (c *Circle) CollideCircle(o *Circle) bool {
	return square(c.x-o.x)+square(c.y-o.y) < square(c.radius+o.radius)
}
