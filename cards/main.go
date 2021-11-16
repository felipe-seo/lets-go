package main

//var card string

func main() {
	//var card string = "Ace of Spades"
	cards := deck{newCard(), newCard()}
	cards = append(cards, "Mais um")
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
