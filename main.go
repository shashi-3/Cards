package main

func main() {
	cards := newdeck()

	hand, remaining := deal(cards, 5)
	hand.print()
	remaining.print()
}
