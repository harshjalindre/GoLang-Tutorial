package main

import (
	"fmt"
	"math"
)

// 1. Define the interface
type Shape interface {
	Area() float64
}

// 2. Define concrete structs
type Circle struct {
	Radius float64
}

type Square struct {
	Side float64
}

// 3. Implement the methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// 4. Use the interface as a parameter
func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	s := Square{Side: 10}

	// Both work because they implement Shape!
	printArea(c)
	printArea(s)
}