package main

import (
	"fmt"
	"strconv"
)

// CPF simula o documento de identificação de pessoa física brasileira
type CPF string

func main() {
	cpf := CPF("11144477735")
	if cpf.EhValido() {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}

// EhValido retorna se um CPF é valido ou não
func (c CPF) EhValido() bool {

	if len(c) != 11 ||
		c == "00000000000" ||
		c == "11111111111" ||
		c == "22222222222" ||
		c == "33333333333" ||
		c == "44444444444" ||
		c == "55555555555" ||
		c == "66666666666" ||
		c == "77777777777" ||
		c == "88888888888" ||
		c == "99999999999" {
		return false
	}

	firstVerificator := 0
	multiplier := 10
	//this loop accumulates a value to calculate the first verification digit
	for i := 0; i < 9; i++ {
		//converts the current digit to string
		currentDigit, err := strconv.Atoi(string(c[i]))
		if err != nil {
			fmt.Println("Conversion error")
		}

		firstVerificator += currentDigit * multiplier
		//fmt.Println(firstVerificator, multiplier, currentDigit)
		multiplier--
	}
	if firstVerificator%11 < 2 {
		firstVerificator = 0
	} else {
		firstVerificator = 11 - (firstVerificator % 11)
	}

	//second verification digit

	secondVerificator := 0
	multiplier = 11

	for i := 0; i < 10; i++ {
		//converts the current digit to string
		currentDigit, err := strconv.Atoi(string(c[i]))
		if err != nil {
			fmt.Println("Conversion error")
		}

		secondVerificator += currentDigit * multiplier
		multiplier--
	}
	if secondVerificator%11 < 2 {
		secondVerificator = 0
	} else {
		secondVerificator = 11 - (secondVerificator % 11)
	}

	//put the tenth and eleventh characters from the input into variables
	userValidationOne, err := strconv.Atoi(string(c[9]))
	if err != nil {
		fmt.Println("Conversion error")
	}
	userValidationTwo, err := strconv.Atoi(string(c[10]))
	if err != nil {
		fmt.Println("Conversion error")
	}

	fmt.Printf("\n Expected values: %v %v \n input values: %v %v\n", firstVerificator, secondVerificator, userValidationOne, userValidationTwo)
	//compare them with the expected values
	if userValidationOne != firstVerificator || userValidationTwo != secondVerificator {
		return false
	} else {
		return true
	}
}
