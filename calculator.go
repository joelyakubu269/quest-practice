package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

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

var Type = map[string][]string{
	"1": {"Circle", "Radius"},
	"2": {"Triangle", "Base", "Height"},
	"3": {"Rectangle", "Length", "Breadth"},
}

func (p Rectangle) Area() float64 {
	return p.Breadth * p.Length
}
func GetUserInput() {
	fmt.Println("welcome Enter the shapes you wish to calculate their Areas")
	for {
		fmt.Println("1: for Circle, 2:for Triangle,3:for rectangle,: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("input cannot be empty")
			continue
		}
		val, ok := Type["input"]
		if ok {
			fmt.Printf("please enter the %v\n: ", val[0])
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

		}
	}
}
func ShapeCal(val []Shape) {}
