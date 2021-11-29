package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Parsons",
		contactInfo: contactInfo{
			email:   "jimp@gmail.com",
			zipCode: 657877,
		},
	}
	jim.updateName("Nicole")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("\n %+v", p)
}
