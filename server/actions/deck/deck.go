package deck

import (
	"fmt"
	"log"

	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

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

	log.Output(1, encounter.ActiveEntity.Id)
	fmt.Printf("%+v\n", originalEntityHandCards)
	fmt.Printf("%+v\n", originalEntityDeckCards)

	updatedEntityHandCards := addCard(originalEntityHandCards, originalEntityDeckCards[0])
	updatedEntityDeckCards := removeCard(originalEntityDeckCards, 1)

	encounter.ActiveEntity.Hand.Cards = updatedEntityHandCards
	encounter.ActiveEntity.Deck.Cards = updatedEntityDeckCards

	log.Output(1, encounter.ActiveEntity.Id)
	fmt.Printf("%+v\n", encounter.ActiveEntity.Hand.Cards)
	fmt.Printf("%+v\n", encounter.ActiveEntity.Deck.Cards)

	return true
}
