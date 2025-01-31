package card

import (
	"fmt"
	"strings"
)

type Card struct {
	s Suit
	r Rank
}

func New(r Rank, s Suit) Card {
	return Card{
		s: s,
		r: r,
	}
}

// erwartet einen rank, returned true wenn
// der rank der karte dem rank im parameter entspricht
func (c Card) IsRank(r Rank) bool {

	return c.r == r
}

// gibt den Rank der karte zurück
func (c Card) GetRank() Rank {

	return c.r

}

// git den Suit der karte zurück
func (c Card) GetSuit() Suit {

	return c.s

}

// String liefert eine AsciiArt-Repräsentation der Karte.
// Zeichnet einen Rahmen um die Karte und innen die Zeichen für
// Rang und Farbe.
// Beispiel Pik Ass:
// ┌───────┐
// │A      │
// │       │
// │   ♠   │
// │       │
// │      A│
// └───────┘
func (c Card) String() string {
	lines := []string{
		"┌───────┐",
		"│%-2s     │",
		"│       │",
		"│   %s   │",
		"│       │",
		"│     %2s│",
		"└───────┘",
	}

	card_string_template := strings.Join(lines, "\n")
	card_string := fmt.Sprintf(card_string_template, c.r, c.s, c.r)
	return card_string
}

// Matches erwartet eine weitere Karte o und prüft,
// ob c und o nach den MauMau-Regeln zusammenpassen.
func (c Card) Matches(o Card) bool {
	// Wenn Farbe oder Zahl gleich sind, passen die
	// Karten zusammen
	return c.r == o.r || c.s == o.s
}
