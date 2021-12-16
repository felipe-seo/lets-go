package main

import (
	"fmt"
	"reflect"
)

type newInt32 struct {
	number      int32
	description string
}

type newFloat32 struct {
	number      float32
	description string
}

type newNumbers interface {
	printType()
}

func printType(nn newNumbers) error {
	nn.printType()

	return fmt.Errorf("wrong type")
}

func main() {
	firstInt := newInt32{2, "int32"}
	firstFloat := newFloat32{2.23, "float32"}
	printType(newInt32(firstInt))
	printType(newFloat32(firstFloat))

}

/*
func (nf newFloat32) printType() {
	fmt.Println(nf.description)
}

func (ni newInt32) printType() {
	fmt.Println(ni.description)
}
*/

//alternative solution that needs no description field in the type
func (nf newFloat32) printType() {
	fmt.Println(reflect.TypeOf(nf.number))
}

func (ni newInt32) printType() {
	fmt.Println(reflect.TypeOf(ni.number))
}
