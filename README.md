# 🃏 Go Playing Cards

A lightweight, modular Go package for creating, shuffling, and managing a standard deck of 52 playing cards. Perfect for building CLI card games like Poker, Blackjack, or Solitaire.

## 🚀 Features

*   **Complete Deck Generation**: Create a standard 52-card deck (4 suits, 13 ranks).
*   **Fisher-Yates Shuffle**: Reliable randomization using Go's `math/rand`.
*   **Dealing Logic**: Simple methods to deal hands and track remaining cards.
*   **File Persistence**: Save your deck state to a local file and load it back later.
*   **Unit Tested**: Includes a robust test suite for deck size and randomization.

## 🛠 Installation

```bash
go get ://github.com

package main

import "fmt"

func main() {
    // 1. Create a new deck
    deck := newDeck()

    // 2. Shuffle the deck
    deck.shuffle()

    // 3. Deal a hand (e.g., 5 cards)
    hand, remainingDeck := deal(deck, 5)

    // 4. See the results
    hand.print()
    fmt.Println("Cards remaining:", len(remainingDeck))
}
```

📂 Project Structure
deck.go: Core logic, types, and receiver functions.

deck_test.go: Automated tests for deck creation and file I/O.

main.go: Entry point for quick demonstrations.
```go
go test
```
