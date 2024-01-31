package main

import (
	"fmt"
	"strings"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}
func (t triangle) double() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}

type ustring string

func (s ustring) Upper() string {
	return strings.ToUpper(string(s))
}

type coloredTriangle struct {
	triangle
	color string
}

func (ct coloredTriangle) perimeter() int {
	return ct.size * 3 * 2
}

func main() {
	t := triangle{3}
	perimeter := t.perimeter()
	fmt.Println(perimeter)

	s := square{5}
	fmt.Println(s.perimeter())

	t.double()
	fmt.Println(t.perimeter())

	str := ustring("go")
	fmt.Println(str, str.Upper())

	ct := coloredTriangle{triangle{5}, "red"}
	fmt.Println(ct.size, ct.perimeter())
}
