package deck

import (
	"cardgames/card"
	"math/rand"
)

// Ein Kartenstapel ist eine Liste von Karten.
type Deck struct {
	Cards []card.Card
}

// NewDeck32 gibt einen neuen Kartenstapel zurück.
func NewDeck32() Deck {
	// Skat-Blatt: 32 Karten, 7-10, Bube, Dame, König, Ass
	cards := []card.Card{
		card.New(card.Seven, card.Clubs),
		card.New(card.Eight, card.Clubs),
		card.New(card.Nine, card.Clubs),
		card.New(card.Ten, card.Clubs),
		card.New(card.Jack, card.Clubs),
		card.New(card.Queen, card.Clubs),
		card.New(card.King, card.Clubs),
		card.New(card.Ace, card.Clubs),
		card.New(card.Seven, card.Spades),
		card.New(card.Eight, card.Spades),
		card.New(card.Nine, card.Spades),
		card.New(card.Ten, card.Spades),
		card.New(card.Jack, card.Spades),
		card.New(card.Queen, card.Spades),
		card.New(card.King, card.Spades),
		card.New(card.Ace, card.Spades),
		card.New(card.Seven, card.Hearts),
		card.New(card.Eight, card.Hearts),
		card.New(card.Nine, card.Hearts),
		card.New(card.Ten, card.Hearts),
		card.New(card.Jack, card.Hearts),
		card.New(card.Queen, card.Hearts),
		card.New(card.King, card.Hearts),
		card.New(card.Ace, card.Hearts),
		card.New(card.Seven, card.Diamonds),
		card.New(card.Eight, card.Diamonds),
		card.New(card.Nine, card.Diamonds),
		card.New(card.Ten, card.Diamonds),
		card.New(card.Jack, card.Diamonds),
		card.New(card.Queen, card.Diamonds),
		card.New(card.King, card.Diamonds),
		card.New(card.Ace, card.Diamonds),
	}

	return Deck{Cards: cards}
}

// Mischt das Deck.
func (d *Deck) Shuffle() {

	shuffled := []card.Card{}

	// ziehe eine random karte und pack sie in eine neue liste

	for len(d.Cards) > 0 {

		shuffled = append(shuffled, d.Draw_Random())

	}

	d.Cards = shuffled

	// wenn das deck leer ist, speicher die neue liste in d.Cards

}

func (d *Deck) Draw_Random() card.Card {

	// einen random index im deck wählen
	index := rand.Intn(len(d.Cards))

	// karte unter diesem index abspeichern
	c := d.Cards[index]

	// index aus der slice löschen

	d.Cards = append(d.Cards[:index], d.Cards[(index+1):]...)

	return c
}

// Zieht eine Karte vom Deck.
// D.h. entfernt die oberste Karte und gibt sie zurück.
func (d *Deck) Draw() card.Card {

	c := d.Cards[0]

	d.Cards = d.Cards[1:]

	return c
}

// Gibt die oberste Karte des Decks zurück.
func (d *Deck) Top() card.Card {

	return d.Cards[0]

}

// Fügt eine Karte zum Deck hinzu.
func (d *Deck) Add(c card.Card) {

	d.Cards = append([]card.Card{c}, d.Cards...)

}

// Gibt die Anzahl der Karten im Deck zurück.
func (d *Deck) Len() int {

	return len(d.Cards)
}
