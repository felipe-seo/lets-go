package main

import "fmt"

type sale struct {
	value       int
	salesperson string
	date        string
}

func main() {
	sale1 := sale{
		value: 15, salesperson: "Sinatra", date: "28/04/2021",
	}

	sale2 := sale{
		value: 10, salesperson: "Scultz", date: "26/04/2021",
	}

	sale3 := sale{
		value: 10, salesperson: "Giordana", date: "27/04/2021",
	}
	sale4 := sale{
		value: 15, salesperson: "Sinatra", date: "28/04/2021",
	}

	sale5 := sale{
		value: 5, salesperson: "Scultz", date: "26/04/2021",
	}

	sale6 := sale{
		value: 10, salesperson: "Giordana", date: "27/04/2021",
	}
	sale7 := sale{
		value: 0, salesperson: "", date: "27/04/2021",
	}
	saleList := []sale{
		sale1, sale2, sale3, sale4, sale5, sale6,
	}

	fmt.Println(calculateSales(saleList))
	fmt.Println(calculateSales2(sale1, sale2, sale4, sale5, sale7))
}

//Might be more useful to receive all the parameters in a single slice
func calculateSales(sl []sale) (map[string]int, int) {
	total := 0
	salesByPerson := make(map[string]int)
	for _, v := range sl {
		salesByPerson[v.salesperson] += v.value
		total += v.value
	}
	return salesByPerson, total
}

//in case there's a need to receive all the parameters separately
func calculateSales2(sl ...sale) (map[string]int, int) {
	total := 0
	salesByPerson := make(map[string]int)
	for _, v := range sl {
		salesByPerson[v.salesperson] += v.value
		total += v.value
	}
	return salesByPerson, total
}
