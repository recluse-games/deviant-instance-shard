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

	turn := &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	entity := &deviant.Entity{
		Deck: GenerateDeckLiteral(5),
		Hand: GenerateHandLiteral(5),
	}

	// Test Draw Turn Phase Action
	if Process(turn, entity, deviant.EntityActionNames_NOTHING) != true {
		t.Fail()
	}

	if len(entity.Hand.Cards) != 6 {
		t.Fail()
	}

	turn = &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_POINT,
	}

	// Test GrantAP Turn Phase Action
	if Process(turn, entity, deviant.EntityActionNames_NOTHING) != true {
		t.Fail()
	}

	if entity.Ap != 5 {
		t.Fail()
	}
}
