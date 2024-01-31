package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

func (s Square) Area() float64 {
	return s.size * s.size
}

type Circle struct {
	radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	squ := Square{5}
	info(squ)

	cir := Circle{5}
	info(cir)
}

func info(s Shape) {
	fmt.Printf("Object: %T\n", s)
	fmt.Println(s, s.Perimeter(), s.Area())
}
