package main

import (
	"fmt"
	"time"
)

func main() {
	//exercise 1
	printGreeting()

	//exercise 2
	printGreeting2()

	//exercise 3
	printGreeting3(27)

	//exercise 3 alternative
	printGreeting3Alt()

	//exercise 3 alternative
	fmt.Println("\n", generateGreeting())

}

//exercise 1
func printGreeting() {
	fmt.Println("1Good morning!!")
}

//exercise 2
func printGreeting2() {
	var name string
	fmt.Println("2Hello, what's your name?")
	fmt.Scanln(&name)
	fmt.Printf("Good morning, %v! \n", name)
}

//exercise 3
func printGreeting3(hour int) {
	var name string

	fmt.Println("\n3Hello, what's your name?")
	fmt.Scanln(&name)
	switch {
	case hour > 0 && hour < 7:
		fmt.Printf("Good early morning, %v!", name)
	case hour > 6 && hour < 12:
		fmt.Printf("Good morning, %v!", name)
	case hour > 12 && hour < 18:
		fmt.Printf("Good afternoon, %v!", name)
	case hour > 18 && hour < 24:
		fmt.Printf("Good evening, %v!", name)
	case hour >= 24 || hour < 0:
		fmt.Printf("It's a weird time, isn't it?")
	}
}

//exercise 3 Alt
func printGreeting3Alt() {
	var name string
	fmt.Println("\n3altHello, what's your name?")
	fmt.Scanln(&name)
	hour := time.Now().Hour()
	switch {
	case hour < 6:
		fmt.Printf("Good early morning, %v!", name)
	case hour > 6 && hour < 12:
		fmt.Printf("Good morning, %v!", name)
	case hour > 12 && hour < 18:
		fmt.Printf("Good afternoon, %v!", name)
	case hour > 18:
		fmt.Printf("Good evening, %v!", name)

	}
}

//exercise 4
func generateGreeting() string {
	hour := time.Now().Hour()
	var greeting string
	switch {
	case hour < 6:
		greeting = "Good early morning"
	case hour > 6 && hour < 12:
		greeting = "Good morning"
	case hour > 12 && hour < 18:
		greeting = "Good afternoon"
	case hour > 18:
		greeting = "Good evening"
	}
	return greeting
}
