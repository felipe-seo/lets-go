package main

import "fmt"

func main() {
	var inputNumber int
	fmt.Println("Enter a number")
	fmt.Scanln(&inputNumber)
	if inputNumber == 0 {
		fmt.Println("The number is zero")
	} else if inputNumber%2 == 0 && inputNumber%3 == 0 {
		fmt.Println("The number is a multiple of 2 and 3")
	} else if inputNumber%2 == 0 {
		fmt.Println("The number is a multiple of 2")
	} else if inputNumber%3 == 0 {
		fmt.Println("The number is a multiple of 3")
	} else {
		fmt.Printf("The number %v is NOT zero, multiple of 2 or 3", inputNumber)
	}

}
