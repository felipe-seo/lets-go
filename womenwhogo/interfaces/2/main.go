package main

import "fmt"

type Animal interface {
	present()
}

type rabbit struct {
	diet string
	filo string
}

type seaAnemone struct {
	diet string
	filo string
}

type lion struct {
	diet string
	filo string
}

func main() {
	rogerRabbit := rabbit{
		diet: "Herbivorous",
		filo: "Chordata",
	}

	meany := seaAnemone{
		diet: "Carnivorous",
		filo: "Cnidaria",
	}

	leo := lion{
		diet: "Carnivorous",
		filo: "Chordata",
	}

	Animal.present(rogerRabbit)
	Animal.present(meany)
	Animal.present(leo)

}

func (r rabbit) present() {
	fmt.Printf("Filo:%v Diet:%v", r.filo, r.diet)
}

func (sa seaAnemone) present() {
	fmt.Printf("Filo:%v Diet:%v", sa.filo, sa.diet)
}

func (l lion) present() {
	fmt.Printf("Filo:%v Diet:%v", l.filo, l.diet)
}
