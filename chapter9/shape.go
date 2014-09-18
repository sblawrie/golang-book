type Shape interface {
    area() float64
    perimeter() float64
}

func (r *Rectangle) perimeter() float64 {
    return r.x1 + r.x2 + r.y1 + r.y2
}

func (c *Circle) perimeter() float64 {
    return math.Pi * c.r * 2
}