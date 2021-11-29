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

func main() {
	t1 := triangle{
		height: 4,
		base:   6,
	}
	sq1 := square{
		sideLength: 10,
	}

	//fmt.Println(t1.getArea())
	//fmt.Println(sq1.getArea())
	printArea(t1)
	printArea(sq1)
}

func printArea(s shape) float64 {
	fmt.Println(s.getArea())
	return 0
}

func (t triangle) getArea() float64 {
	return t.base * t.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
