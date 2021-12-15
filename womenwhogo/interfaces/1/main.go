package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type geometricArea interface {
	getArea() float64
}

func main() {
	var sq1 square
	sq1.side = 10

	var circle1 circle
	circle1.radius = 2

	showArea(sq1)
	showArea(circle1)
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (c circle) getArea() float64 {
	return c.radius * c.radius * 3.141592
}

func showArea(g geometricArea) {
	fmt.Printf("The area is: %v", g.getArea())
}
