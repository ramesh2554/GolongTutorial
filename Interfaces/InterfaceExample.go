package main

import (
	"fmt"
	"math"
)

/*
In this example we find a area and perimeter of area using interface


*/

// declare a interface
type geometry interface {
	area() float64
	perimeter() float64
}

// create rectangle struct to find area, perimeter of rectangle and circle

type rectangle struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {

	return 2 * (r.length + r.width)
}
func (c circle) area() float64 {
	return 2 * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {

	fmt.Println("Area  --->", g.area())
	fmt.Println("perimeter  ----> ", g.perimeter())

}
func main() {
	r := rectangle{length: 3, width: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
