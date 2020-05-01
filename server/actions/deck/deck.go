package deck

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

func addCard(slice []*deviant.Card, card *deviant.Card) []*deviant.Card {
	return append(slice, card)
}

func removeCard(slice []*deviant.Card, s int) []*deviant.Card {
	if len(slice) > 1 {
		return append(slice[:s], slice[s+1:]...)
	}

	return nil
}

// DrawCard draws a card for entity.
func DrawCard(entity *deviant.Entity) bool {
	originalEntityDeckCards := entity.Deck.Cards
	originalEntityHandCards := entity.Hand.Cards

	updatedEntityHandCards := addCard(originalEntityHandCards, originalEntityDeckCards[0])
	updatedEntityDeckCards := removeCard(originalEntityDeckCards, 1)

	entity.Hand.Cards = updatedEntityHandCards
	entity.Deck.Cards = updatedEntityDeckCards

	return true
}
