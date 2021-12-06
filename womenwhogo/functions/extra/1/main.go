package main

import "fmt"

func main() {
	numberList := []int{1, 2, 3, 4}
	doubleOddsHalfEvens(numberList)
	fmt.Println(numberList)
}

//exercise 1

func doubleOddsHalfEvens(ns []int) ([]int, int) {
	var total int
	//resultSlice := ns
	for i, v := range ns {
		if v%2 == 0 {
			ns[i] = v / 2
		} else {
			ns[i] = v + v
		}

		fmt.Println(i, v)
		total += v
	}
	fmt.Println(ns, total)
	return ns, total
}
