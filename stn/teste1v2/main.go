package main

import (
	"errors"
	"fmt"
)

type item struct {
	name   string
	amount int
	price  int
}

func main() {
	//crio os items
	item1 := item{
		name:   "batata",
		amount: 1,
		price:  100,
	}
	item2 := item{
		name:   "tomate",
		amount: 1,
		price:  109,
	}
	//crio a lista de items
	itemList := map[int]item{
		1: item1,
		2: item2,
	}
	//printo a lista
	fmt.Println("Lista de items: ", itemList)

	//crio a lista de e-mails
	emailList := []string{
		"mary@mail.com",
		"john@mail.com",
		"rose@mail.com",
		"sunny@mail.com",
		"julia@mail.com",
		"debra@mail.com",
		"adam@mail.com",
	}
	//printo a lista
	fmt.Println("Lista de emails: ", emailList)
	divideTab(itemList, emailList)
}

//função que calcula a conta, recebe os itens e os clientes
func divideTab(il map[int]item, el []string) (map[string]int, error) {
	total := 0
	tab := make(map[string]int, len(el))

	if verifyItems(il) {
		return tab, errors.New("itemsError")
	}
	if verifyEmails(el) {
		return tab, errors.New("emailsError")
	}
	//fazer condição para interromper caso venha alguma lista defeituosa

	for _, element := range il {
		total += element.price * element.amount
	}
	fmt.Println("Total: ", total)
	remainder := int(total) % len(el)
	fmt.Println("Resto: ", remainder)
	totalPerAccount := total / len(el)
	fmt.Println("Divisão: ", totalPerAccount)
	//distribuir a sobra, centavo por centavo, até não ter mais sobra
	for _, v := range el {
		if remainder > 0 {
			tab[v] = totalPerAccount + 1
			remainder--
		} else {
			tab[v] = totalPerAccount
		}
	}
	fmt.Println("Conta: ", tab)
	return tab, errors.New("error")
}
func verifyItems(il map[int]item) bool {
	if len(il) == 0 {
		fmt.Println("Lista vazia")
		return true
	}
	for _, v := range il {
		if v.name == "" || v.amount == 0 || v.price == 0 {
			fmt.Printf("Algo de errado com o item: %v", v)
			return true
		}
	}
	return false
}

func verifyEmails(el []string) bool {
	if len(el) == 0 {
		fmt.Println("Lista vazia")
		return true
	}
	//fazer algo para verificar se tem já tem o email na lista
	for _, v := range el {
		if v == "" {
			fmt.Printf("Algo de errado com o email: %v", v)
			return true
		}
	}
	return false
}
