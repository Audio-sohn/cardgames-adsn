package hand

import (
	"cardgames/card"
	"fmt"
	"slices"
)

type Hand struct {
	Cards []card.Card
}

// New gibt eine leere Hand zurück.
func New() Hand {
	return Hand{}
}

// Add fügt eine Karte zur Hand hinzu.
func (h *Hand) Add(c card.Card) {
	h.Cards = append(h.Cards, c)
}

// String gibt eine AsciiArt-Repräsentation der Hand zurück.
func (h Hand) String() string {

	// string to build

	hand_string := ""

	// strings for universal building blocks
	card_top := "┌───────┐"

	card_upper_template := "│%-2s     │"

	card_middle_template := "│   %s   │"

	card_lower_template := "│     %2s│"

	card_space := "│       │"
	card_bottom := "└───────┘"

	// number of cards in the hand
	// card_count := len(h.Cards)

	// build Hand image line by line
	// every card is 7 lines high
	for i := 0; i < 7; i++ {

		// iterate through each card per line
		for j, speci := range h.Cards {

			switch i {

			case 0:
				hand_string += fmt.Sprintf(card_top)

			case 1:
				hand_string += fmt.Sprintf(card_upper_template, speci.GetRank())

			case 2:
				hand_string += fmt.Sprintf(card_space)

			case 3:
				hand_string += fmt.Sprintf(card_middle_template, speci.GetSuit())

			case 4:
				hand_string += fmt.Sprintf(card_space)

			case 5:
				hand_string += fmt.Sprintf(card_lower_template, speci.GetRank())

			case 6:
				hand_string += fmt.Sprintf(card_bottom)

			}

			// add whitespace inbetween cards, but not at the end
			if j < h.Len()-1 {
				hand_string += " "
			}

		}

		// add newline at the end of each line
		hand_string += "\n"
	}

	return hand_string
}

// Remove entfernt eine Karte aus der Hand.
func (h *Hand) Remove(c card.Card) bool {

	// check if c is on the hand, save position
	index := slices.Index(h.Cards, c)

	// if not, return false
	if index == -1 {

		return false

		//	if yes, remove the entry
	} else {

		h.Cards = append(h.Cards[:index], h.Cards[(index+1):]...)
	}

	return true
}

// Len gibt die Anzahl der Karten in der Hand zurück.
func (h Hand) Len() int {

	return len(h.Cards)
}

// ContainsRank prüft, ob ein bestimmter Rang in der Hand ist.
func (h Hand) ContainsRank(r card.Rank) bool {

	// durch hand iterieren
	for _, speci := range h.Cards {

		// prüfen ob rank der karte == r
		if speci.GetRank() == r {

			return true

		}
	}

	// wenn nicht gefunden, false returnen
	return false
}
