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
		Ap:    0,
		MaxAp: 5,
		Deck:  GenerateDeckLiteral(5),
		Hand:  GenerateHandLiteral(5),
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entity,
		Turn:         turn,
		Board: &deviant.Board{
			Tiles: &deviant.Tiles{
				Tiles: []*deviant.TilesRow{
					{
						Tiles: []*deviant.Tile{
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
						},
					},
				},
			},
			OverlayTiles: []*deviant.Tile{},
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

	// Test Draw Turn Phase Action
	if Process(encounter, deviant.EntityActionNames_NOTHING, nil, nil) != true {
		t.Fail()
	}

	if len(entity.Hand.Cards) != 6 {
		t.Fail()
	}

	turn = &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_POINT,
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entity,
		Turn:         turn,
		Board: &deviant.Board{
			Tiles: &deviant.Tiles{
				Tiles: []*deviant.TilesRow{
					{
						Tiles: []*deviant.Tile{
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
							{Id: "grass_0000"},
						},
					},
				},
			},
			OverlayTiles: []*deviant.Tile{},
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

	// Test GrantAP Turn Phase Action
	if Process(encounter, deviant.EntityActionNames_NOTHING, nil, nil) != true {
		t.Fail()
	}

	if encounter.ActiveEntity.Ap != 5 {
		t.Logf("%v", encounter.ActiveEntity)
		t.Fail()
	}
}
