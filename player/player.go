package player

// define the Player interface
// the interface is exposed

type Player interface {

	// make a move
	GetMove()

	// show the current Hand
	GetHand()

	// draw a new card
	DrawCard()
}
