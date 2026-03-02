package main

func main() {
	//	cards := newdeck()
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
}
