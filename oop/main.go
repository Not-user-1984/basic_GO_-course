package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Площадь: %f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circ := Circle{Radius: 2}

	printArea(rect)
	printArea(circ)
}
