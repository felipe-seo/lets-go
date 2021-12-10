package main

import "fmt"

type broccoli int

var x broccoli
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T ", x)
	x = 911
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
}
