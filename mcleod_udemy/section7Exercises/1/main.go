package main

import "fmt"

func main() {
	//exercise1
	inputNumber := 0
	fmt.Println("Enter a number: ")
	fmt.Scanln(&inputNumber)
	fmt.Printf("The number %v, is %#b in binary, %#o in octal, and %#x in hexadecimal ", inputNumber, inputNumber, inputNumber, inputNumber)

	//exercise2
	a := 2 == 1+1
	b := 2 >= 2
	c := 2 <= 3
	d := 3 < 6
	e := 1 > 0
	f := 1 != 10
	fmt.Println(a, b, c, d, e, f)

	//exercise3
	pi := 3.1415
	var myConstant int = 63

	fmt.Println(pi, myConstant)
}
