package deck

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

func addCard(slice []*deviant.Card, card *deviant.Card) []*deviant.Card {
	return append(slice, card)
}

func removeCard(slice []*deviant.Card, s int) []*deviant.Card {
	if len(slice) > 1 {
		return append(slice[:s], slice[s+1:]...)
	}

	return []*deviant.Card{}
}

// DrawCard draws a card for entity.
func DrawCard(encounter *deviant.Encounter) bool {
	originalEntityDeckCards := encounter.ActiveEntity.Deck.Cards
	originalEntityHandCards := encounter.ActiveEntity.Hand.Cards

	updatedEntityHandCards := addCard(originalEntityHandCards, originalEntityDeckCards[0])
	updatedEntityDeckCards := removeCard(originalEntityDeckCards, 1)

	encounter.ActiveEntity.Hand.Cards = updatedEntityHandCards
	encounter.ActiveEntity.Deck.Cards = updatedEntityDeckCards

	return true
}
