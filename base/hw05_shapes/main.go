package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height
}

func calculateArea(s any) (float64, error) {
	if shape, ok := s.(Shape); ok {
		return shape.Area(), nil
	}
	return 0, errors.New("not found shape")
}

func main() {
	circle := Circle{5}
	rectangle := Rectangle{5, 3}
	triangle := Triangle{1, 1}
	invalid := "not a figure"

	shapes := []any{circle, rectangle, triangle, invalid}
	for _, shape := range shapes {
		area, err := calculateArea(shape)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			switch s := shape.(type) {
			case Circle:
				fmt.Printf("Circle: Radius %.0f Area: %f\n", s.Radius, area)
			case Rectangle:
				fmt.Printf("Rectangle: Radius %.0f Area: %f\n", s.Width, s.Height)
			case Triangle:
				fmt.Printf("Triangle: Radius %.0f Area: %f\n", s.Base, s.Height)

			}

		}

	}

}
