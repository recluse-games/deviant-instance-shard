package deck

import (
	"fmt"
	"log"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func removeEntityFromOrder(entityID string, slice []string) []string {
	var entityIDIndex = find(slice, entityID)

	log.Output(1, fmt.Sprintf("%d", entityIDIndex))

	if len(slice) > entityIDIndex+1 {
		return slice[:entityIDIndex+copy(slice[entityIDIndex:], slice[entityIDIndex+1:])]
	} else if len(slice) == entityIDIndex {
		return slice[:entityIDIndex+copy(slice[entityIDIndex:], slice[entityIDIndex-1:])]
	}
	if 0 > entityIDIndex-1 {
		return slice[:0+copy(slice[(entityIDIndex-1):], slice[entityIDIndex:])]
	} else if len(slice) > 1 {
		return slice[:len(slice)-1]
	} else {
		return []string{}
	}
}

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
