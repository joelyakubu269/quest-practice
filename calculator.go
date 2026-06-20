package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var shapes = []Shape{}
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
func GetUserInput() float64 {
	fmt.Println("welcome Enter the shapes you wish to calculate their Areas")
	for {
		fmt.Println("1: for Circle, 2:for Triangle,3:for rectangle,: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("input cannot be empty")
			continue
		}
		val, ok := Type[input]
		if ok {
			switch val[0] {
			case "Circle":
				fmt.Printf("please enter the %v\n: ", val[1])
				num, err := ConvertInput()
				if err != nil {
					continue
				}
				AddShape(Circle{Radius: num})
			case "Triangle":
				fmt.Printf("please enter the %v\n: ", val[1])
				num, err := ConvertInput()
				if err != nil {
					continue
				}
				fmt.Printf("please enter the %v\n: ", val[2])
				num2, err := ConvertInput()
				if err != nil {
					continue
				}
				AddShape(Triangle{Base: num, Height: num2})
			case "Rectangle":
				fmt.Printf("please enter the %v\n: ", val[1])
				num, err := ConvertInput()
				if err != nil {
					continue
				}
				fmt.Printf("please enter the %v\n: ", val[2])
				num2, err := ConvertInput()
				if err != nil {
					continue
				}
				AddShape(Rectangle{Length: num, Breadth: num2})
			}

		}
		fmt.Println("Do you wish to continue (yes/no): ")
		input, _ = reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "no" {
			return ShapeCal(shapes)
		}
	}
}
func ShapeCal(val []Shape) float64 {
	totalArea := 0.0
	for _, s := range val {
		totalArea += s.Area()
	}
	return totalArea
}

func ConvertInput() (float64, error) {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
func AddShape(s Shape) {
	shapes = append(shapes, s)
}
