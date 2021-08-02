package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	s := square{sideLength: 12}
	tri := triangle{height: 12, base: 4}
	printArea(s)
	printArea(tri)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (tri triangle) getArea() float64 {
	return tri.height * tri.base * 0.5
}
