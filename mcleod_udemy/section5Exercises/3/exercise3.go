package main

import "fmt"

var x = 42
var y = "Tomato Bond"
var z = true

func main() {
	s := fmt.Sprintf("x: %v,  y: %v, and z: %v", x, y, z)
	fmt.Println(s)
}
