/*
1) Descubra por que não compila
The type int8 cannot contain such a large number(150).
2) Erros de compilação nos ajudam a compreender o que precisamos consertar em nosso código. O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
The constant overflows our variable, which means 150 is too big for int 8.
3) Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida na tela
*/

package main

import (
	"fmt"
)

func main() {
	var quilometros int16
	quilometros = 150
	fmt.Println(quilometros)
}
