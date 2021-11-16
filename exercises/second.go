package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Bó testar o package de tempo vendo a data e hora atuais")
	date_time := time.Now()
	fmt.Println("Agora são: ", time.Now().Format("15:04:05"))
	fmt.Println("Do dia: ", time.Now().Format("01-02-2007"))
	fmt.Println(date_time.Format("01-02-2006 15:04:05"))

}
