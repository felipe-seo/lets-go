package main

import "fmt"

func main() {
	//fmt.Println(aToC("/abcdefghijklmnopqrstuvwxyz!"))
	fmt.Println(cToA("Aabcdefghijklmnopqrstuvwxyz"))
	//fmt.Println(cToA("abcdefgh"))
}

func aToC(s string) string {
	var newString []byte
	for _, v := range s {
		//x,y,z
		if v > 87 && v <= 90 || v > 119 && v <= 122 {
			newString = append(newString, byte(v-23))
			//a-w
		} else if v > 64 && v <= 97 || v > 89 && v <= 122 {
			newString = append(newString, byte(v+3))
		} /*else {
			newString = append(newString, byte(v))
		}*/ //suppose you want to keep the other symbols

	}
	return (string(newString))
}

func cToA(s string) string {
	var newString []byte
	for _, v := range s {

		//x,y,z
		if v > 64 && v <= 67 || v > 96 && v <= 99 {
			newString = append(newString, byte(v+23))
			fmt.Println(string(v))
			//a-w
		} else if v > 64 && v <= 97 || v > 89 && v <= 122 {
			newString = append(newString, byte(v-3))
			fmt.Println(v)
		} /*else {
			newString = append(newString, byte(v))
		} */ //suppose you want to keep the other symbols

	}
	return (string(newString))
}
