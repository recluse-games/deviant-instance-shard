package deck

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

func addCard(slice []*deviant.Card, card *deviant.Card) []*deviant.Card {
	slice = append(slice, card)
	return slice
}

func removeCard(slice []*deviant.Card) []*deviant.Card {
	return slice[:0+copy(slice[0:], slice[1:])]
}

// DrawCard draws a card for entity.
func DrawCard(encounter *deviant.Encounter) bool {
	var topCard = encounter.ActiveEntity.Deck.Cards[0]

	encounter.ActiveEntity.Deck.Cards = removeCard(encounter.ActiveEntity.Deck.Cards)
	encounter.ActiveEntity.Hand.Cards = addCard(encounter.ActiveEntity.Hand.Cards, topCard)

	return true
}
