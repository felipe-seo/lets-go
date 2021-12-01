package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Write your text and hit enter")
	newText := readLine()

	fmt.Printf("%T", newText)
	countLetters(newText)
}

func readLine() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func countLetters(s string l map[string]int) {
	container := strings.Split(s, "")
	for i, v := range container {
		fmt.Printf("\n", i, v)
		
	}

}
