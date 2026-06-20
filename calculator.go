package main

import (
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}

func (p Circle) Area() float64 {
	return math.Pi * 2 * p.Radius * p.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (p Triangle) Area() float64 {
	return 0.5 * p.Base * p.Height
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (p Rectangle) Area() float64 {
	return p.Breadth * p.Length
}
func ShapeCal(val []Shape)
