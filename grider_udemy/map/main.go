package main

import "fmt"

func main() {
	//colors := make(map[string]string)
	//var colors map[string]string*

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#745000",
		"white": "#ffffff",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
