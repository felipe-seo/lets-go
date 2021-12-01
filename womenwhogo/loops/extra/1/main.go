package main

import "fmt"

func main() {
	loopExtra1()
}

func loopExtra1() {
	littleSlice := []string{
		"Tomatoes", "Pasta", "Rice", "Beans", "Spinach", "Lettuce", "Eggs",
	}
	for i := 0; i < len(littleSlice); i++ {
		fmt.Println(i+1, littleSlice[i])
	}
}
