package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Print("Type the year you were born in: ")
	var birthYear string
	fmt.Scanln(&birthYear)
	currentYear := time.Now().Year()
	age := currentYear - strconv.Atoi(birthYear)
	fmt.Println(age)
}
