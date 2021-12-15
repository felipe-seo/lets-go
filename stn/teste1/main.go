package main

import "fmt"

type item struct {
	name   string
	amount float64
	price  float64
}

func main() {
	//crio os items
	item1 := item{
		name:   "batata",
		amount: 2,
		price:  1.5,
	}
	item2 := item{
		name:   "tomate",
		amount: 4,
		price:  2,
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
	}
	//printo a lista
	fmt.Println("Lista de emails: ", emailList)
	divideTab(itemList, emailList)
}

func divideTab(il map[int]item, el []string) {

	//fmt.Println(il, el)
	//fazer condição para interromper caso venha alguma lista vazia
	total := 0.0
	for _, element := range il {
		total += element.price * element.amount
		fmt.Println(total)
	}
	remainder := int(total) % int(len(el))
	fmt.Println(remainder)
	totalPerAccount := total / float64(len(el))
	fmt.Println(totalPerAccount)
}
