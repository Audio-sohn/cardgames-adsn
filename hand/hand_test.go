package hand

import (
	"cardgames/card"
	"fmt"
)

func ExampleHand_Add() {
	h := New()
	h.Add(card.New(card.Seven, card.Diamonds))
	h.Add(card.New(card.Eight, card.Clubs))
	h.Add(card.New(card.King, card.Hearts))

	fmt.Println(h)

	// Output:
	// ┌───────┐ ┌───────┐ ┌───────┐
	// │7      │ │8      │ │K      │
	// │       │ │       │ │       │
	// │   ♦   │ │   ♣   │ │   ♥   │
	// │       │ │       │ │       │
	// │      7│ │      8│ │      K│
	// └───────┘ └───────┘ └───────┘
}

func ExampleHand_Remove() {

	h := New()
	c := h.Cards[0]
	h.Remove(c)
	c = h.Cards[0]
	h.Remove(c)
	fmt.Println(len(h.Cards))

	//	Output
	//	2
	//	1
}
