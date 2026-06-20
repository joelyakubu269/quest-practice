package main

import "math"

type Shape interface {
	Area() float64
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

type Perimeter interface {
	Perimeter() float64
}

func PrintAreas(shapes []Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}
func PrintPerimeters(shapes []Perimeter) float64 {
	totalPerimeter := 0.0
	for _, s := range shapes {
		totalPerimeter += s.Perimeter()
	}
	return totalPerimeter

}
