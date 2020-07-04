package actions

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
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
	// WARNING: If we don't have cards left, deal damage to myself and then return true.
	if len(encounter.ActiveEntity.Deck.Cards) == 0 {
		encounter.ActiveEntity.Hp = encounter.ActiveEntity.Hp - 1
		return true
	}

	var topCard = encounter.ActiveEntity.Deck.Cards[0]

	encounter.ActiveEntity.Deck.Cards = removeCard(encounter.ActiveEntity.Deck.Cards)
	encounter.ActiveEntity.Hand.Cards = addCard(encounter.ActiveEntity.Hand.Cards, topCard)

	return true
}
