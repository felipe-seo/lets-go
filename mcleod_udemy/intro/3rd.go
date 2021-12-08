package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
