package main

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
func PrintPerimeters(shapes []Shape)
