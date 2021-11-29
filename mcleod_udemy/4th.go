package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Primeira mensagem!")
	//duration := time.Second * 20
	time.Sleep(time.Second * 2)
	fmt.Println("Segunda mensagem")
}
