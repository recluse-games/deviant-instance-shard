package matchmaker

import (
	"github.com/google/uuid"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func generateHandLiterals(size int32) *deviant.Hand {
	deckLiteral := &deviant.Hand{
		Id:    "grass_0000",
		Cards: generateCardLiterals(size),
	}

	return deckLiteral
}

func generateDiscardLiteral(size int32) *deviant.Discard {
	deckLiteral := &deviant.Discard{
		Id:    "grass_0000",
		Cards: generateCardLiterals(size),
	}

	return deckLiteral
}

func generateDeckLiterals(size int32) *deviant.Deck {
	deckLiteral := &deviant.Deck{
		Id:    "grass_0000",
		Cards: generateCardLiterals(size),
	}

	return deckLiteral
}

func generateCardLiterals(size int32) []*deviant.Card {
	var cardLiterals = []*deviant.Card{}

	for i := int32(0); i < size; i++ {
		card := &deviant.Card{
			Id:          "attack_slash_0000",
			BackId:      "back_0000",
			InstanceId:  uuid.New().String(),
			Cost:        0,
			Title:       "Test Title",
			Flavor:      "Test Flavor",
			Description: "Test Description",
			Type:        deviant.CardType_ATTACK,
			Action: &deviant.CardAction{
				Id: uuid.New().String(),
				Pattern: []*deviant.Pattern{
					{
						Direction: deviant.Direction_DOWN,
						Distance:  3,
						Offset: []*deviant.Offset{
							{
								Direction: deviant.Direction_DOWN,
								Distance:  1,
							},
						},
					},
					{
						Direction: deviant.Direction_DOWN,
						Distance:  3,
						Offset: []*deviant.Offset{
							{
								Direction: deviant.Direction_LEFT,
								Distance:  1,
							},
							{
								Direction: deviant.Direction_DOWN,
								Distance:  1,
							},
						},
					},
					{
						Direction: deviant.Direction_DOWN,
						Distance:  3,
						Offset: []*deviant.Offset{
							{
								Direction: deviant.Direction_RIGHT,
								Distance:  1,
							},
							{
								Direction: deviant.Direction_DOWN,
								Distance:  1,
							},
						},
					},
				},
			},
		}

		cardLiterals = append(cardLiterals, card)
	}

	return cardLiterals
}

// GenerateMatch Generates a new match
func GenerateMatch() *deviant.EncounterResponse {
	test := &deviant.EncounterResponse{
		PlayerId: "player_0000",
		Encounter: &deviant.Encounter{
			Id:        "encounter_0000",
			Completed: false,
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
				Entities: &deviant.Entities{
					Entities: []*deviant.EntitiesRow{
						{
							Entities: []*deviant.Entity{
								{
									Id:         "0001",
									Hp:         10,
									Ap:         5,
									Alignment:  deviant.Alignment_FRIENDLY,
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(2),
									Discard:    generateDiscardLiteral(2),
									Initiative: 5,
									OwnerId:    "0001",
								},
								{
									Id:         "0002",
									Hp:         10,
									Ap:         4,
									Alignment:  deviant.Alignment_FRIENDLY,
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(2),
									Discard:    generateDiscardLiteral(0),
									Initiative: 5,
									OwnerId:    "0002",
								},
								{},
								{},
								{},
								{},
								{},
								{},
								{},
							},
						},
						{
							Entities: []*deviant.Entity{
								{},
								{},
								{},
								{},
								{},
								{},
								{},
								{},
							},
						},
						{
							Entities: []*deviant.Entity{
								{},
								{},
								{},
								{},
								{},
								{},
								{},
								{},
							},
						},
						{
							Entities: []*deviant.Entity{
								{},
								{},
								{},
								{},
								{},
								{},
								{},
								{},
							},
						},
						{
							Entities: []*deviant.Entity{
								{},
								{},
								{},
								{},
								{},
								{},
								{},
								{},
							},
						},
						{
							Entities: []*deviant.Entity{
								{},
								{},
								{},
								{},
								{},
								{},
								{},
								{},
							},
						},
						{
							Entities: []*deviant.Entity{
								{},
								{},
								{},
								{},
								{},
								{},
								{},
								{},
							},
						},
						{
							Entities: []*deviant.Entity{
								{
									Id:         "0004",
									Hp:         10,
									Ap:         5,
									Alignment:  deviant.Alignment_UNFRIENDLY,
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(2),
									Discard:    generateDiscardLiteral(0),
									Initiative: 3,
									OwnerId:    "0004",
								},
								{
									Id:         "0005",
									Hp:         10,
									Ap:         5,
									Alignment:  deviant.Alignment_UNFRIENDLY,
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(2),
									Discard:    generateDiscardLiteral(0),
									Initiative: 3,
									OwnerId:    "0005",
								},
								{
									Id:         "0006",
									Hp:         10,
									Ap:         5,
									Alignment:  deviant.Alignment_UNFRIENDLY,
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(2),
									Discard:    generateDiscardLiteral(0),
									Initiative: 3,
									OwnerId:    "0006",
								},
								{},
								{},
								{},
								{},
								{},
							},
						},
					},
				},
			},
			ActiveEntity: &deviant.Entity{
				Id:         "0001",
				Hp:         10,
				Ap:         5,
				Alignment:  deviant.Alignment_FRIENDLY,
				Class:      deviant.Classes_WARRIOR,
				Hand:       generateHandLiterals(0),
				Deck:       generateDeckLiterals(2),
				Discard:    generateDiscardLiteral(0),
				Initiative: 5,
				OwnerId:    "0001",
			},
			ActiveEntityOrder: []string{"0001", "0002", "0003", "0004", "0005", "0006"},
			Turn: &deviant.Turn{
				Id:    "turn_0000",
				Phase: deviant.TurnPhaseNames_PHASE_POINT,
			},
		},
	}

	return test
}
