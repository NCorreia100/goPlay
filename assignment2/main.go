package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func printArea(s shape) {
	fmt.Printf("\n Area of %#v", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	aTriangle := triangle{
		base:   2,
		height: 3.1,
	}

	aSquare := square{
		sideLength: 2.5,
	}

	printArea(aTriangle)
	printArea(aSquare)
}
