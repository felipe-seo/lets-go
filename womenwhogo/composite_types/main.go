package main

import "fmt"

func main() {
	//exercise 1
	array1 := []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh"}
	fmt.Printf(" Type: %T\n Elements: %v", array1, array1)

	//exercise 2
	sliceA := []float64{1.1, 2.2, 3.3}
	sliceB := sliceA
	fmt.Println("\n", sliceB, sliceA)

	//exercise 3
	sl1 := []int{1, 2, 3, 4, 5, 6}
	sl2 := []int{2, 4, 6, 8, 10, 12}
	sl3 := append(sl1, sl2...)
	fmt.Println("\n", sl3)

	//exercise 4
	list1 := []string{"meat", "carrots", "turnip", "milk", "chicken breast", "eggs"}
	fmt.Println(list1)
	list1 = append(list1, "strawberry")
	fmt.Println(list1)

	//exercise 5
	colorMap := map[string]string{
		"orange": "FFA500",
		"white":  "FFFFFF",
		"black":  "000000",
	}
	fmt.Println(colorMap)
	delete(colorMap, "orange")
	fmt.Println(colorMap)

	//exercise 6
	yearMap := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	fmt.Printf("Mês 1: %v \nMês 12: %v", yearMap[1], yearMap[12])
	fmt.Println("\n Map size: ", len(yearMap))

	//exercise 7
	noboru := person{
		name:          "Noboru",
		age:           29,
		favoriteColor: "Orange",
	}

	scrooge := person{
		name:          "Scrooge",
		age:           75,
		favoriteColor: "Gold",
	}

	maria := person{
		name:          "Maria",
		age:           20,
		favoriteColor: "Green",
	}
	fmt.Printf("%v's favorite color is %v and age is %v \n", noboru.name, noboru.favoriteColor, noboru.age)
	fmt.Printf("%v's favorite color is %v and age is %v \n", scrooge.name, scrooge.favoriteColor, scrooge.age)
	fmt.Printf("%v's favorite color is %v and age is %v \n", maria.name, maria.favoriteColor, maria.age)
}

//exercise 7
type person struct {
	name          string
	age           int
	favoriteColor string
}
