package actions

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

func addCard(slice []*deviant.Card, card *deviant.Card) []*deviant.Card {
	slice = append(slice, card)
	return slice
}

func removeCard(slice []*deviant.Card) []*deviant.Card {
	return slice[:0+copy(slice[0:], slice[1:])]
}

// DrawCard Draws a card for the currently active entity of an encounter.
func DrawCard(encounter *deviant.Encounter, logger *zap.Logger) bool {
	if encounter.ActiveEntity.Deck == nil {
		encounter.ActiveEntity.Hp = encounter.ActiveEntity.Hp - 1

		if encounter.ActiveEntity.Hp <= 0 {
			encounter.ActiveEntityOrder = removeEntityFromOrder(encounter.ActiveEntity.Id, encounter.ActiveEntityOrder)
		}

		return true
	}

	// WARNING: I don't love that this is included in the DrawCard logic, it feels like it should be somewhere else.
	if len(encounter.ActiveEntity.Deck.Cards) == 0 {
		encounter.ActiveEntity.Hp = encounter.ActiveEntity.Hp - 1

		if encounter.ActiveEntity.Hp <= 0 {
			encounter.ActiveEntityOrder = removeEntityFromOrder(encounter.ActiveEntity.Id, encounter.ActiveEntityOrder)
		}

		return true
	}

	topCard := encounter.ActiveEntity.Deck.Cards[0]

	encounter.ActiveEntity.Deck.Cards = removeCard(encounter.ActiveEntity.Deck.Cards)
	encounter.ActiveEntity.Hand.Cards = addCard(encounter.ActiveEntity.Hand.Cards, topCard)

	if logger != nil {
		logger.Debug("Card Drawn",
			zap.String("actionID", "DrawCard"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return true
}
