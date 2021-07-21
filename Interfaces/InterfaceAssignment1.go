package main

import "fmt"

type shape interface {
	printArea() float64
}
type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

func (sq square) printArea() float64 {
	return sq.sideLength * sq.sideLength
}
func (tri triangle) printArea() float64 {
	return 0.5 * tri.base * tri.height
}

func calculate(s shape) {
	fmt.Println(s.printArea())
}
func main() {
	square := square{5}
	tri := triangle{3, 4}
	calculate(square)
	calculate(tri)
}
