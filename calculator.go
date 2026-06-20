package main

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
func GetUserInput() {
	fmt.Println("welcome Enter the shapes you wish to calculate their Areas")
	fmt.Println("1: for Circle, 2:for Triangle,3:for rectangle,")
	fmt.Println("you may enter more than one number if you wish")
}
func ShapeCal(val []Shape) {}
