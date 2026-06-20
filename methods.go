package main

import (
	"fmt"
	"math"
)

type Shape interface {
	//Area() float64
	Perimeter() float64
}
type Triangle struct {
	Base   float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (p Triangle) Area() float64 {
	return p.Base * p.Height
}
func (p Triangle) Perimeter() float64 {
	return 2 * (p.Base + p.Height)
}
func (p Circle) Area() float64 {
	return math.Pi * p.Radius * p.Radius
}
func (p Circle) Perimeter() float64 {
	return math.Pi * 2 * p.Radius
}

func PrintAreas(shapes []Shape) (float64, float64) {
	totalArea := 0.0
	totalPerimeter := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
		totalPerimeter += s.Perimeter()
	}

	return totalArea, totalPerimeter
}

func main() {
	shapes := []Shape{
		Triangle{Base: 20, Height: 30},
		Circle{Radius: 25},
	}

	fmt.Println(PrintAreas(shapes))

}
