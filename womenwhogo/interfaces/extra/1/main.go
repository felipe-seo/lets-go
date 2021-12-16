package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type printer interface {
	printContent()
}

type newString struct {
	txt string
}

type newFile struct {
	target *os.File
}

func main() {
	var myFile newFile
	myFile.target, _ = os.OpenFile("index.html", os.O_RDONLY, os.ModeDevice)
	//printContent(myFile)

	myString := newString{"this is my content"}
	//printContent(myString)
	printer.printContent(myFile)
	printer.printContent(myString)
}

func (nf newFile) printContent() {
	content, _ := ioutil.ReadAll(nf.target)
	fmt.Printf("%s", string(content))
}

func (ns newString) printContent() {
	fmt.Printf("%v", ns)
}
