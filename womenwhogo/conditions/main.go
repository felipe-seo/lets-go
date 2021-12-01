package main

import (
	"fmt"
	"time"
)

func main() {
	//exercise 1
	var birthYear int
	currentYear := time.Now().Year()
	fmt.Println("Type your birthyear")
	fmt.Scanln(&birthYear)
	age := currentYear - birthYear
	if age < 18 {
		fmt.Printf("Underage, cannot vote if you are younger than 18\n")
	} else {
		fmt.Printf("You can vote\n")
	}

	//exercise 2
	someNumber := 18
	if someNumber > 0 {
		fmt.Println("Positive")
	} else if someNumber < 0 {
		fmt.Println("Negative")
	}

	//exercise 3
	fmt.Println(ageCategory(age))

	//exercise 4
	fmt.Println(ageCategorySwitch(age))

	//exercise 5
	hour := 22

	fmt.Println(periodDay(hour))
	fmt.Println(periodDay(200))

}

//exercise 3
func ageCategory(age int) string {
	if age < 18 {
		return "Underage"
	} else if age > 60 {
		return "Elderly"
	} else {
		return "Adult"
	}

}

//exercise 4
func ageCategorySwitch(age int) string {
	var category string
	switch {
	case (age < 18):
		category = "Underage"
	case (age >= 18 && age <= 60):
		category = "Adult"
	case (age > 60):
		category = "Elderly"
	}
	return category
}

//exercise 5
func periodDay(t int) string {
	var period string
	switch {
	case (t >= 0 && t < 6):
		period = "Early morning"
	case (t >= 6 && t < 12):
		period = "Morning"
	case (t >= 12 && t < 6):
		period = "Early morning"
	case (t >= 12 && t < 18):
		period = "Afternoon"
	case (t >= 18 && t <= 23):
		period = "Night"
	case (t < 0 || t > 23):
		period = "Value out of range, please enter a number between 0 and 23"
	}
	return period
}
