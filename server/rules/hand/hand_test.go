package hand

import (
	"strconv"
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

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

func TestValidateSize(t *testing.T) {

	testEntityWithValidHand := &deviant.Entity{
		Hand: &deviant.Hand{
			Id:    "0000",
			Cards: GenerateCardLiterals(1),
		},
	}

	isHandSizeValid := ValidateSize(testEntityWithValidHand)

	if isHandSizeValid != true {
		t.Fail()
	}

	testEntityWithInvalidHand := &deviant.Entity{
		Hand: &deviant.Hand{
			Id:    "0000",
			Cards: GenerateCardLiterals(20),
		},
	}

	isHandSizeInvalid := ValidateSize(testEntityWithInvalidHand)

	if isHandSizeInvalid != false {
		t.Fail()
	}
}
