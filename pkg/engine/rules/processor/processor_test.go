package processor

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

func TestProcess(t *testing.T) {

	entityWithEnoughCardsInDeckToDraw := &deviant.Entity{
		Deck: GenerateDeckLiteral(5),
		Hand: GenerateHandLiteral(5),
	}

	turn := &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entityWithEnoughCardsInDeckToDraw,
		Turn:         turn,
	}

	// Test Draw Action in Correct Phase
	if Process(encounter, deviant.EntityActionNames_DRAW, nil, nil) != true {
		t.Fail()
	}

	turn = &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_ACTION,
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entityWithEnoughCardsInDeckToDraw,
		Turn:         turn,
	}

	// Test Draw Action in Incorrect Phase
	if Process(encounter, deviant.EntityActionNames_DRAW, nil, nil) == false {
		t.Logf("%s", "Failed attempting to process")

		t.Fail()
	}

	entityWithoutEnoughCardsInDeckToDraw := &deviant.Entity{
		Deck: GenerateDeckLiteral(0),
		Hand: GenerateHandLiteral(5),
	}

	turn = &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entityWithoutEnoughCardsInDeckToDraw,
		Turn:         turn,
	}

	if Process(encounter, deviant.EntityActionNames_DRAW, nil, nil) != false {
		t.Logf("%s", "Failed attempting to process")
		t.Fail()
	}

}
