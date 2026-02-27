package main

import "fmt"

func main() {
	cards := newdeck()
	fmt.Println(cards.toString())
}
