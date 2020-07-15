package actions

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap"
)

//add Adds a card to the end of a provided slice
func add(slice []*deviant.Card, card *deviant.Card) []*deviant.Card {
	slice = append(slice, card)
	return slice
}

//draw Removes a card object from the front of a provided slice
func draw(slice []*deviant.Card) []*deviant.Card {
	return slice[:0+copy(slice[0:], slice[1:])]
}

// DrawCard Draws a card for the currently active entity of an encounter.
func DrawCard(encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	if encounter.ActiveEntity.Deck == nil {
		encounter.ActiveEntity.Hp = encounter.ActiveEntity.Hp - 1
		return true
	} else if len(encounter.ActiveEntity.Deck.Cards) == 0 {
		encounter.ActiveEntity.Hp = encounter.ActiveEntity.Hp - 1
		return true
	}

	topCard := encounter.ActiveEntity.Deck.Cards[0]

	encounter.ActiveEntity.Deck.Cards = draw(encounter.ActiveEntity.Deck.Cards)
	encounter.ActiveEntity.Hand.Cards = add(encounter.ActiveEntity.Hand.Cards, topCard)

	logger.Debug("Card Drawn",
		zap.String("actionID", "DrawCard"),
		zap.String("entityID", encounter.ActiveEntity.Id),
	)

	return true
}
