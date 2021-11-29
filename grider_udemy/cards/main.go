package main

//var card string

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
