package main

import "fmt"

func main() {
	testString := "aaaAA BB CccC"
	fmt.Println(letterCounter(testString, "c"))
	//fmt.Println(letterCounter("Go is awesome", "g"))
}

func letterCounter(str string, letter string) int {
	letterOccurences := 0

	//fmt.Printf("\nO tipo de %v é: %T\n", letter, letter)
	//fmt.Println(letter)

	//fmt.Printf("\n String: %v", str)
	for _, v := range str {
		if string(v) == letter || string(v+32) == letter {
			fmt.Println(string(v + 32))
			letterOccurences++
		}
	}
	return letterOccurences
}
