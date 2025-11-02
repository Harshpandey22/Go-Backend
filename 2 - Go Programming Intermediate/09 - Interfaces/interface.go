package main

import (
	"fmt"
	"math"
)

// An interface specifies what a type can do, not what its internal data structure is. It separates the behavior from the implementation.
// Use: The core uses of interfaces in Go are to enable polymorphism and decoupling, which are essential for building flexible, testable, and maintainable software.

// Syntax
// type InterfaceName interface {
// 	MethodOne()
// 	MethodTwo(parameter string) int
// 	MethodThree() (result int, err error)
// }

// Interface with blank functions
type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) perim() float64 {
	return 2 * (r.width * r.height)
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := Rect{width: 3, height: 4} // Creating Instance of Rect
	c := Circle{radius: 5}         // Creating Instance of Circle

	measure(r)
	measure(c)

	// Output:
	// 	{3 4}
	// 12
	// 24
	// {5}
	// 78.53981633974483
	// 31.41592653589793
}
