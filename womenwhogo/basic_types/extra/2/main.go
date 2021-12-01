package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Type the year you were born in: ")
	var birthYearString int
	fmt.Scanln(&birthYearString)
	currentYear := time.Now().Year()
	//birthYear, e := strconv.Atoi(birthYearString)
	//if e == nil {
	//	fmt.Printf("%T \n %v", birthYear, birthYear)
	//}

	age := currentYear - birthYearString
	fmt.Println("\n", age)
}
