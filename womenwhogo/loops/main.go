package main

import "fmt"

func main() {
	//exercise 1
	//loopExercise1()
	//exercise 2
	//loopExercise2()
	//exercise 3
	//loopExercise3()
	//loopExercise4()
	//loopExercise5()
	loopExercise6()
}

/*
func loopExercise1() {
	for i := 13; i <= 27; i++ {
		fmt.Printf("\n%v", i)
	}
}

func loopExercise2() {
	i := 0
	for i <= 23 {
		fmt.Printf("\n%v", i)
		i++
	}
}

func loopExercise3() {
	i := 0
	for i <= 23 {
		m := 0
		for m <= 59 {
			fmt.Printf("\n%vh%vm", i, m)
			m++
		}
		i++
	}
}

func loopExercise4() {
	i := 0
	for i <= 2 {
		m := 0
		for m <= 59 {
			s := 0
			for s <= 59 {
				fmt.Printf("\n%vh%vm%v", i, m, s)
				s++
			}
			m++
		}
		i++
	}
}

func loopExercise5() {
	littleSlice := []string{
		"zero", "one", "two", "three", "four", "five",
	}
	for i, v := range littleSlice {
		fmt.Println(v, i)
	}
}
*/
func loopExercise6() {
	littleSlice := []string{
		"Tomatoes", "Pasta", "Rice", "Beans", "Spinach", "Lettuce", "Eggs",
	}
	for i, v := range littleSlice {
		fmt.Println(i+1, v)
	}
}
