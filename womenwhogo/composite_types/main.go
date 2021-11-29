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
}
