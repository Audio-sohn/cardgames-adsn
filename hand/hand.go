package hand

import (
	"cardgames/card"
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
	// TODO
	return ""
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
