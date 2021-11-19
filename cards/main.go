package main

//var card string

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
}
