package enginetest

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
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

func generateMatchHandLiterals(size int32, class deviant.Classes) *deviant.Hand {
	deckLiteral := &deviant.Hand{
		Id:    "grass_0000",
		Cards: generateCardLiterals(size, class),
	}

	return deckLiteral
}

func generateMatchDiscardLiterals(size int32, class deviant.Classes) *deviant.Discard {
	deckLiteral := &deviant.Discard{
		Id:    "grass_0000",
		Cards: generateCardLiterals(size, class),
	}

	return deckLiteral
}

func generateMatchDeckLiterals(size int32, class deviant.Classes) *deviant.Deck {
	deckLiteral := &deviant.Deck{
		Id:    "grass_0000",
		Cards: generateCardLiterals(size, class),
	}

	return deckLiteral
}

func generateCardLiterals(size int32, class deviant.Classes) []*deviant.Card {
	var cardLiterals = []*deviant.Card{}

	switch class {
	case deviant.Classes_WARRIOR:
		for i := int32(0); i < size; i++ {
			card := &deviant.Card{
				Id:          "attack_bash_0000",
				BackId:      "back_0000",
				InstanceId:  uuid.New().String(),
				Cost:        3,
				Damage:      2,
				Title:       "Bash",
				Flavor:      "OP Area Move",
				Description: "Something Too Broken to Be Real",
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
							Distance:  1,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_LEFT,
									Distance:  1,
								},
								{
									Direction: deviant.Direction_DOWN,
									Distance:  3,
								},
							},
						},
						{
							Direction: deviant.Direction_DOWN,
							Distance:  1,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_RIGHT,
									Distance:  1,
								},
								{
									Direction: deviant.Direction_DOWN,
									Distance:  3,
								},
							},
						},
					},
				},
			}

			cardLiterals = append(cardLiterals, card)

			card = &deviant.Card{
				Id:          "attack_slash_0000",
				BackId:      "back_0000",
				InstanceId:  uuid.New().String(),
				Cost:        2,
				Damage:      2,
				Title:       "Slash",
				Flavor:      "Downward Dog",
				Description: "A Simple Slash",
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
					},
				},
			}

			cardLiterals = append(cardLiterals, card)
		}
	case deviant.Classes_PRIEST:
		for i := int32(0); i < size; i++ {
			card := &deviant.Card{
				Id:          "cast_heal_0000",
				BackId:      "back_0000",
				InstanceId:  uuid.New().String(),
				Cost:        1,
				Damage:      2,
				Title:       "Heal",
				Flavor:      "A Basic Heal",
				Description: "A Simple Heal",
				Type:        deviant.CardType_HEAL,
				Action: &deviant.CardAction{
					Id: uuid.New().String(),
					Pattern: []*deviant.Pattern{
						{
							Direction: deviant.Direction_DOWN,
							Distance:  1,
							Offset: []*deviant.Offset{
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

			card = &deviant.Card{
				Id:          "cast_self_heal_0000",
				BackId:      "back_0000",
				InstanceId:  uuid.New().String(),
				Cost:        2,
				Damage:      2,
				Title:       "Bandage",
				Flavor:      "A Basic Self Heal",
				Description: "A Simple Self Heal",
				Type:        deviant.CardType_HEAL,
				Action: &deviant.CardAction{
					Id: uuid.New().String(),
					Pattern: []*deviant.Pattern{
						{
							Direction: deviant.Direction_DOWN,
							Distance:  1,
							Offset:    []*deviant.Offset{},
						},
					},
				},
			}

			cardLiterals = append(cardLiterals, card)

			card = &deviant.Card{
				Id:          "cast_healing_ray_0000",
				BackId:      "back_0000",
				InstanceId:  uuid.New().String(),
				Cost:        3,
				Damage:      2,
				Title:       "Healing Ray",
				Flavor:      "A Basic Heal",
				Description: "A Ranged Heal",
				Type:        deviant.CardType_HEAL,
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
					},
				},
			}

			cardLiterals = append(cardLiterals, card)

			card = &deviant.Card{
				Id:          "cast_healing_lob_0000",
				BackId:      "back_0000",
				InstanceId:  uuid.New().String(),
				Cost:        2,
				Damage:      1,
				Title:       "Healing Lob",
				Flavor:      "A Basic Heal",
				Description: "A Ranged Lob Heal",
				Type:        deviant.CardType_HEAL,
				Action: &deviant.CardAction{
					Id: uuid.New().String(),
					Pattern: []*deviant.Pattern{
						{
							Direction: deviant.Direction_DOWN,
							Distance:  1,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_DOWN,
									Distance:  3,
								},
							},
						},
					},
				},
			}

			cardLiterals = append(cardLiterals, card)
		}
	case deviant.Classes_MAGE:
		for i := int32(0); i < size; i++ {
			card := &deviant.Card{
				Id:          "attack_fireball_0000",
				BackId:      "back_0000",
				InstanceId:  uuid.New().String(),
				Cost:        2,
				Damage:      2,
				Title:       "Fireball",
				Flavor:      "Dunking",
				Description: "A Simple Fireball",
				Type:        deviant.CardType_ATTACK,
				Action: &deviant.CardAction{
					Id: uuid.New().String(),
					Pattern: []*deviant.Pattern{
						{
							Direction: deviant.Direction_DOWN,
							Distance:  1,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_DOWN,
									Distance:  3,
								},
							},
						},
					},
				},
			}

			cardLiterals = append(cardLiterals, card)

			card = &deviant.Card{
				Id:          "attack_searing_touch_0000",
				BackId:      "back_0000",
				InstanceId:  uuid.New().String(),
				Cost:        2,
				Damage:      3,
				Title:       "Searing Touch",
				Flavor:      "Burn Baby",
				Description: "A Cross Attack",
				Type:        deviant.CardType_ATTACK,
				Action: &deviant.CardAction{
					Id: uuid.New().String(),
					Pattern: []*deviant.Pattern{
						{
							Direction: deviant.Direction_RIGHT,
							Distance:  1,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_UP,
									Distance:  2,
								},
								{
									Direction: deviant.Direction_RIGHT,
									Distance:  2,
								},
							},
						},
						{
							Direction: deviant.Direction_RIGHT,
							Distance:  1,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_DOWN,
									Distance:  2,
								},
								{
									Direction: deviant.Direction_RIGHT,
									Distance:  2,
								},
							},
						},
						{
							Direction: deviant.Direction_LEFT,
							Distance:  1,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_UP,
									Distance:  2,
								},
								{
									Direction: deviant.Direction_LEFT,
									Distance:  2,
								},
							},
						},
						{
							Direction: deviant.Direction_LEFT,
							Distance:  1,
							Offset: []*deviant.Offset{
								{
									Direction: deviant.Direction_DOWN,
									Distance:  2,
								},
								{
									Direction: deviant.Direction_LEFT,
									Distance:  2,
								},
							},
						},
					},
				},
			}

			cardLiterals = append(cardLiterals, card)
		}
	}

	// Suffle the Deck
	rand.Seed(time.Now().UnixNano())
	for i := len(cardLiterals) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		cardLiterals[i], cardLiterals[j] = cardLiterals[j], cardLiterals[i]
	}

	return cardLiterals
}

//GenerateMatchObject Generates a mock match object.
func GenerateMatchObject() *deviant.EncounterResponse {
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
								{
									Id:         "0001",
									Name:       "Ian",
									Hp:         10,
									MaxHp:      10,
									Ap:         5,
									MaxAp:      5,
									Alignment:  deviant.Alignment_FRIENDLY,
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateMatchHandLiterals(0, deviant.Classes_WARRIOR),
									Deck:       generateMatchDeckLiterals(10, deviant.Classes_WARRIOR),
									Discard:    generateMatchDiscardLiterals(0, deviant.Classes_WARRIOR),
									Initiative: 5,
									OwnerId:    "0001",
									Rotation:   deviant.EntityRotationNames_WEST,
								},
								{},
								{},
								{
									Id:         "0003",
									Name:       "Zach",
									Hp:         10,
									MaxHp:      10,
									Ap:         5,
									MaxAp:      5,
									Alignment:  deviant.Alignment_FRIENDLY,
									Class:      deviant.Classes_MAGE,
									Hand:       generateMatchHandLiterals(0, deviant.Classes_MAGE),
									Deck:       generateMatchDeckLiterals(10, deviant.Classes_MAGE),
									Discard:    generateMatchDiscardLiterals(0, deviant.Classes_MAGE),
									Initiative: 5,
									OwnerId:    "0001",
									Rotation:   deviant.EntityRotationNames_NORTH,
								},
								{
									Id:         "0005",
									Name:       "Chris",
									Hp:         10,
									MaxHp:      10,
									Ap:         5,
									MaxAp:      5,
									Alignment:  deviant.Alignment_FRIENDLY,
									Class:      deviant.Classes_PRIEST,
									Hand:       generateMatchHandLiterals(0, deviant.Classes_PRIEST),
									Deck:       generateMatchDeckLiterals(10, deviant.Classes_PRIEST),
									Discard:    generateMatchDiscardLiterals(0, deviant.Classes_PRIEST),
									Initiative: 5,
									OwnerId:    "0001",
									Rotation:   deviant.EntityRotationNames_NORTH,
								},
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
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        2,
									MaxHp:     2,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        2,
									MaxHp:     2,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        2,
									MaxHp:     2,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        2,
									MaxHp:     2,
									State:     deviant.EntityStateNames_IDLE,
									Class:     deviant.Classes_WALL,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        2,
									MaxHp:     2,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        2,
									MaxHp:     2,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        2,
									MaxHp:     2,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        2,
									MaxHp:     2,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
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
									Id:         "0002",
									Name:       "Cameron",
									Hp:         10,
									MaxHp:      10,
									Ap:         5,
									MaxAp:      5,
									Alignment:  deviant.Alignment_UNFRIENDLY,
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateMatchHandLiterals(0, deviant.Classes_WARRIOR),
									Deck:       generateMatchDeckLiterals(10, deviant.Classes_WARRIOR),
									Discard:    generateMatchDiscardLiterals(0, deviant.Classes_WARRIOR),
									Initiative: 5,
									OwnerId:    "0002",
									Rotation:   deviant.EntityRotationNames_SOUTH,
								},
								{},
								{
									Id:         "0004",
									Name:       "Matt",
									Hp:         10,
									MaxHp:      10,
									Ap:         5,
									MaxAp:      5,
									Alignment:  deviant.Alignment_UNFRIENDLY,
									Class:      deviant.Classes_MAGE,
									Hand:       generateMatchHandLiterals(0, deviant.Classes_MAGE),
									Deck:       generateMatchDeckLiterals(10, deviant.Classes_MAGE),
									Discard:    generateMatchDiscardLiterals(0, deviant.Classes_MAGE),
									Initiative: 5,
									OwnerId:    "0002",
									Rotation:   deviant.EntityRotationNames_SOUTH,
								},
								{
									Id:         "0006",
									Name:       "Ben",
									Hp:         10,
									MaxHp:      10,
									Ap:         5,
									MaxAp:      5,
									Alignment:  deviant.Alignment_UNFRIENDLY,
									Class:      deviant.Classes_PRIEST,
									Hand:       generateMatchHandLiterals(0, deviant.Classes_PRIEST),
									Deck:       generateMatchDeckLiterals(10, deviant.Classes_PRIEST),
									Discard:    generateMatchDiscardLiterals(0, deviant.Classes_PRIEST),
									Initiative: 5,
									OwnerId:    "0002",
									Rotation:   deviant.EntityRotationNames_SOUTH,
								},
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
				Name:       "Ian",
				Hp:         10,
				MaxHp:      10,
				Ap:         5,
				MaxAp:      5,
				Alignment:  deviant.Alignment_FRIENDLY,
				Class:      deviant.Classes_WARRIOR,
				Hand:       generateMatchHandLiterals(0, deviant.Classes_WARRIOR),
				Deck:       generateMatchDeckLiterals(10, deviant.Classes_WARRIOR),
				Discard:    generateMatchDiscardLiterals(0, deviant.Classes_WARRIOR),
				Initiative: 5,
				OwnerId:    "0001",
				Rotation:   deviant.EntityRotationNames_WEST,
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
