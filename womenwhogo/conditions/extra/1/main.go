package main

import "fmt"

func main() {
	var a, b, c int
	a = 1500
	b = 1500
	c = 3000
	if a-b >= 0 && a-c >= 0 {
		fmt.Println(a)
	} else if b-a >= 0 && b-c >= 0 {
		fmt.Println(b)
	} else if c-a >= 0 && c-b >= 0 {
		fmt.Println(c)
	}

}
