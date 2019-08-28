package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.length)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{length: 3, width: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
