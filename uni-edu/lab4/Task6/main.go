package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func (r *Rectangle) Scale(k float64) {
	r.Width *= k
	r.Height *= k
}

type Circle struct{ Radius float64 }

func (c *Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c *Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

var _ Shape = Rectangle{}
var _ Shape = (*Rectangle)(nil)
var _ Shape = (*Circle)(nil)

func main() {
	r := Rectangle{3, 4}
	r.Scale(2)
	fmt.Printf("after Scale(2): %+v\n", r) 

	shapes := []Shape{
		Rectangle{3, 4},
		&Circle{Radius: 1},
	}
	for _, s := range shapes {
		fmt.Printf("%T: area=%.2f perimeter=%.2f\n", s, s.Area(), s.Perimeter())
	}