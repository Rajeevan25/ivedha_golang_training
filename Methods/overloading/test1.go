package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}
type Rectangle struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func calculate(s Shape) {
	fmt.Println("Area of the shape ", s.area())
	fmt.Println("Perimeter of the shape ", s.perimeter())
}
func main() {
	shapes := []Shape{
		Circle{radius: 7},
		Rectangle{length: 4, width: 5},
	}
	for _, v := range shapes {
		calculate(v)
	}

}
