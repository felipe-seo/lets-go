package main

import "fmt"

type condo struct {
	number int
	owner  string
	garage bool
}

func main() {
	expensiveCondo := condo{
		number: 100,
		owner:  "Clay Guida",
		garage: true,
	}

	cheapCondo := condo{
		number: 5,
		owner:  "Homer Simpson",
		garage: false,
	}
	condoSlice := []condo{
		expensiveCondo,
		cheapCondo,
	}

	for i := range condoSlice {
		if condoSlice[i].garage == true {
			fmt.Printf("Condo number %v has a garage and is owned by %v\n", condoSlice[i].number, condoSlice[i].owner)
		} else {
			fmt.Printf("Condo number %v does not have a garage and is owned by %v\n", condoSlice[i].number, condoSlice[i].owner)
		}
	}
	//fmt.Println(condoSlice)
}
