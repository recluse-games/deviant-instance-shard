package deck

import (
	"strconv"
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func GenerateHandLiteral(size int32) *deviant.Hand {
	deckLiteral := &deviant.Hand{
		Id:    "0000",
		Cards: GenerateCardLiterals(size),
	}

	return deckLiteral
}

func GenerateDeckLiteral(size int32) *deviant.Deck {
	deckLiteral := &deviant.Deck{
		Id:    "0000",
		Cards: GenerateCardLiterals(size),
	}

	return deckLiteral
}

func GenerateCardLiterals(size int32) []*deviant.Card {
	var cardLiterals = []*deviant.Card{}

	for i := int32(0); i < size; i++ {
		card := &deviant.Card{
			Id:          strconv.FormatInt(int64(i), 10),
			Cost:        0,
			Title:       "Test Title",
			Flavor:      "Test Flavor",
			Description: "Test Description",
			Type:        deviant.CardType_ATTACK,
			Action: &deviant.CardAction{
				Id: "0000",
				Pattern: []*deviant.Pattern{
					{
						Direction: deviant.Direction_DOWN,
						Distance:  0,
					},
				},
			},
		}

		cardLiterals = append(cardLiterals, card)
	}

	return cardLiterals
}

func TestDrawCard(t *testing.T) {
	entityWithOneCardInDeckAndHand := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Deck: GenerateDeckLiteral(1),
			Hand: GenerateHandLiteral(1),
		},
	}

	if DrawCard(entityWithOneCardInDeckAndHand) != true {
		t.Fail()
	}

	if len(entityWithOneCardInDeckAndHand.ActiveEntity.Deck.Cards) != 0 {
		t.Fail()
	}

	if len(entityWithOneCardInDeckAndHand.ActiveEntity.Hand.Cards) != 2 {
		t.Fail()
	}

	entityWithTwoCardsInDeckAndNoneInHand := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Deck: GenerateDeckLiteral(2),
			Hand: GenerateHandLiteral(0),
		},
	}

	if DrawCard(entityWithTwoCardsInDeckAndNoneInHand) != true {
		t.Fail()
	}

	if len(entityWithTwoCardsInDeckAndNoneInHand.ActiveEntity.Deck.Cards) != 1 {
		t.Fail()
	}

	if len(entityWithTwoCardsInDeckAndNoneInHand.ActiveEntity.Hand.Cards) != 1 {
		t.Fail()
	}

	if DrawCard(entityWithTwoCardsInDeckAndNoneInHand) != true {
		t.Fail()
	}

	if len(entityWithTwoCardsInDeckAndNoneInHand.ActiveEntity.Deck.Cards) != 0 {
		t.Fail()
	}

	if len(entityWithTwoCardsInDeckAndNoneInHand.ActiveEntity.Hand.Cards) != 2 {
		t.Fail()
	}
}
