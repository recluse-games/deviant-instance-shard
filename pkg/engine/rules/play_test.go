package rules

import (
	"testing"

	"github.com/google/uuid"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap/zaptest"
)

func TestValidatePlayApCost(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Hp: 5,
		Ap: 5,
		Hand: &deviant.Hand{
			Cards: []*deviant.Card{
				{
					Id:          "0001",
					InstanceId:  "12345",
					Cost:        4,
					Title:       "Test Title",
					Flavor:      "Test Flavor",
					Description: "Test Description",
					Type:        deviant.CardType_ATTACK,
					Actions: []*deviant.CardAction{
						{
							Id:     uuid.New().String(),
							Origin: true,
							Status: deviant.CardActionStatusTypes_EMPTY,
						},
					},
				},
			},
		},
	}

	playAction := &deviant.EntityPlayAction{
		Id:     "0001",
		CardId: "12345",
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entity,
	}

	result := ValidatePlayApCost(entity, playAction, encounter, logger.Sugar())

	if result != true {
		t.Fail()
	}
}

func TestValidateCardInHand(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Hp: 5,
		Ap: 5,
		Hand: &deviant.Hand{
			Cards: []*deviant.Card{
				{
					Id:          "0001",
					InstanceId:  "12345",
					Cost:        4,
					Title:       "Test Title",
					Flavor:      "Test Flavor",
					Description: "Test Description",
					Type:        deviant.CardType_ATTACK,
					Actions: []*deviant.CardAction{
						{
							Id:     uuid.New().String(),
							Origin: true,
							Status: deviant.CardActionStatusTypes_EMPTY,
						},
					},
				},
			},
		},
	}

	playAction := &deviant.EntityPlayAction{
		Id:     "0001",
		CardId: "12345",
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entity,
	}

	result := ValidateCardInHand(entity, playAction, encounter, logger.Sugar())
	if result != true {
		t.Logf("Failed to validate that card that should be in hand.")
		t.Fail()
	}

	playAction = &deviant.EntityPlayAction{
		Id:     "0001",
		CardId: "1111",
	}

	result = ValidateCardInHand(entity, playAction, encounter, logger.Sugar())
	if result != false {
		t.Logf("Failed to detect card that shouldn't be in hand")
		t.Fail()
	}

}

func TestValidateCardConstraints(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Hp: 5,
		Ap: 5,
		Hand: &deviant.Hand{
			Cards: []*deviant.Card{
				{
					Id:          "0001",
					InstanceId:  "12345",
					Cost:        4,
					Title:       "Test Title",
					Flavor:      "Test Flavor",
					Description: "Test Description",
					Type:        deviant.CardType_ATTACK,
					Actions: []*deviant.CardAction{
						{
							Id:     uuid.New().String(),
							Origin: true,
							Status: deviant.CardActionStatusTypes_EMPTY,
						},
					},
				},
			},
		},
	}

	playAction := &deviant.EntityPlayAction{
		Id:     "0001",
		CardId: "12345",
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entity,
		Board: &deviant.Board{
			Entities: &deviant.Entities{
				Entities: []*deviant.EntitiesRow{
					{
						Entities: []*deviant.Entity{
							entity,
						},
					},
				},
			},
		},
	}

	result := ValidateCardConstraints(entity, playAction, encounter, logger.Sugar())
	if result != true {
		t.Logf("Failed to validate that card that should be in hand.")
		t.Fail()
	}
}
