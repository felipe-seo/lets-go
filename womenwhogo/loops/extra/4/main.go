package main

import "fmt"

func main() {
	userInput := 1
	for userInput%2 != 0 {
		fmt.Println("Type an even number")
		fmt.Scanln(&userInput)
	}
	fmt.Println("Bingo! Now we can share it equally")
}
