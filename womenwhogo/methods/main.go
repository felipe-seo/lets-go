package main

import "fmt"

type circle struct {
	radius float64
}

type intSlice struct {
	numbers []int
}

type intSlice2 []int

func main() {
	//exercise1
	circle1 := circle{
		1.5,
	}
	circle1.Area()
	circle1.Perimeter()

	//exercise2
	myIntSlice := intSlice{
		numbers: []int{1, 2, 3},
	}

	myIntSlice.Average()

}

//exercise1
func (c circle) Area() float64 {
	result := (3.1415 * c.radius * c.radius)
	fmt.Printf("\nThe area of the circle is: %.3v", result)
	return result
}
func (c circle) Perimeter() float64 {
	result := (c.radius * 2 * 3.1415)
	fmt.Printf("\nThe perimeter of the circle is: %.3v", result)
	return result
}

//exercise2
func (s intSlice) Average() float64 {
	total := 0.0
	for _, v := range s.numbers {
		total += float64(v)
	}
	total = total / float64(len(s.numbers))
	fmt.Println("\nThe average is: ", total)
	return total
}
