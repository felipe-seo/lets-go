package main

import "fmt"

func main() {
	//exercise 1
	var numA int
	var numB int
	var numC int
	numA = 1
	numB = 2
	numC = 3
	fmt.Println(numA, "(exercise 1)")
	fmt.Println(numB, "(exercise 1)")
	fmt.Println(numC, "(exercise 1)")
	//exercise 2
	a := 230
	b := 27
	sumAB := a + b
	subAB := a - b
	fmt.Printf("\n %v \n %v \n %v \n %v", a, b, sumAB, subAB)

	//exercise 3
	bananaPrice := 1.25
	cervejaPrice := 2.98
	abacatePrice := 4.59
	salgadinhoPrice := 7.29

	bananaAmount := 2.170
	cervejaAmount := 6.0
	abacateAmount := 5.650
	salgadinhoAmount := 3.0

	bananaTotal := bananaPrice * bananaAmount
	cervejaTotal := cervejaPrice * cervejaAmount
	abacateTotal := abacatePrice * abacateAmount
	salgadinhoTotal := salgadinhoPrice * salgadinhoAmount
	total := bananaTotal + cervejaTotal + abacateTotal + salgadinhoTotal

	fmt.Printf("Banana: %.2f \n Cerveja: %.2f \n Abacate: %.2f \n Salgadinho: %.2f \n Total a pagar: %.2f", bananaTotal, cervejaTotal, abacateTotal, salgadinhoTotal, total)

	//exercise 4

}
