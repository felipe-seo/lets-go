package main

import "fmt"

//Create a new type called deck
//It is a slice(variable size array) of strings

type deck []string

//receiver sets up methods on variables that we create
//receiver on a function \/ inside the parenthesis
//the d stands for the variable
//that will come and it's of the type deck, in this
//case it's the variable cards
//in other words it's a reference to the actual
//variable of the type deck that we're working with atm
func (d deck) print() {
	///\ actual copy of the variable
	for i, card := range d {
		fmt.Println(i, card)
	}
}
