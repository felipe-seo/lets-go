package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//Open the file and attribute its value or the error to the variables on the left
	myFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//open the file and show it on standard output
	fmt.Println("--------------- Start of file ---------------")
	io.Copy(os.Stdout, myFile)
	fmt.Println("--------------- End of file ---------------")
}
