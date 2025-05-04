package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius int
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.Radius*c.Radius)
}

type Rectangle struct {
	Length int
	Width  int
}

func (r *Rectangle) Area() float64 {
	return float64(r.Length * r.Width)
}

func ShapeFactory(shapeType string, params ...float64) Shape {
	switch shapeType {
	case "circle":
		if len(params) >= 1 {
			return &Circle{
				Radius: int(params[0]),
			}
		}
	case "rectangle":
		if len(params) >= 2 {
			return &Rectangle{
				Length: int(params[0]),
				Width:  int(params[1]),
			}
		}
	}
	return nil
}

func main() {
	circle := ShapeFactory("circle", 6)
	if c, ok := circle.(*Circle); ok {
		fmt.Printf("Circle Radius = %d, Area = %.2f\n", c.Radius, c.Area())
	}
	//fmt.Printf("Circle with area = %.2f\n", circle.Area())

	rectangle := ShapeFactory("rectangle", 4, 6)
	if r, ok := rectangle.(*Rectangle); ok {
		fmt.Printf("Rectangle Length = %d, Width = %d, Area = %.2f\n", r.Length, r.Width, r.Area())
	}
	//fmt.Printf("Rectangle with area = %.2f\n", rectangle.Area())
}
