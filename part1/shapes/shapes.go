package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func ShapeSlice(sl []Shape) {
	for _, v := range sl {
		fmt.Printf("Area: %f\n", v.Area())
	}
}
