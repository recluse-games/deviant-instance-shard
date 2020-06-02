package matchmaker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewCacheClient Get a new cache client
func NewCacheClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

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

		card = &deviant.Card{
			Id:          "cast_heal_0000",
			BackId:      "back_0000",
			InstanceId:  uuid.New().String(),
			Cost:        2,
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

		card = &deviant.Card{
			Id:          "cast_healing_ray_0000",
			BackId:      "back_0000",
			InstanceId:  uuid.New().String(),
			Cost:        2,
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
			Id:          "attack_fireball_0000",
			BackId:      "back_0000",
			InstanceId:  uuid.New().String(),
			Cost:        1,
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
	}

	// Suffle the Deck
	rand.Seed(time.Now().UnixNano())
	for i := len(cardLiterals) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		cardLiterals[i], cardLiterals[j] = cardLiterals[j], cardLiterals[i]
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
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(10),
									Discard:    generateDiscardLiteral(0),
									Initiative: 5,
									OwnerId:    "0001",
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
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(10),
									Discard:    generateDiscardLiteral(0),
									Initiative: 5,
									OwnerId:    "0001",
								},
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
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        1,
									MaxHp:     1,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        1,
									MaxHp:     1,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        1,
									MaxHp:     1,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        1,
									MaxHp:     1,
									State:     deviant.EntityStateNames_IDLE,
									Class:     deviant.Classes_WALL,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        1,
									MaxHp:     1,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        1,
									MaxHp:     1,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        1,
									MaxHp:     1,
									Class:     deviant.Classes_WALL,
									State:     deviant.EntityStateNames_IDLE,
									Alignment: deviant.Alignment_NEUTRAL,
								},
								{
									Id:        uuid.New().String(),
									Name:      "Wall",
									Hp:        1,
									MaxHp:     1,
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
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(10),
									Discard:    generateDiscardLiteral(0),
									Initiative: 5,
									OwnerId:    "0001",
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
									Class:      deviant.Classes_WARRIOR,
									Hand:       generateHandLiterals(0),
									Deck:       generateDeckLiterals(10),
									Discard:    generateDiscardLiteral(0),
									Initiative: 5,
									OwnerId:    "0001",
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
				Name:       "Ian",
				Hp:         10,
				MaxHp:      10,
				Ap:         5,
				MaxAp:      5,
				Alignment:  deviant.Alignment_FRIENDLY,
				Class:      deviant.Classes_WARRIOR,
				Hand:       generateHandLiterals(0),
				Deck:       generateDeckLiterals(10),
				Discard:    generateDiscardLiteral(0),
				Initiative: 5,
				OwnerId:    "0001",
			},
			ActiveEntityOrder: []string{"0001", "0002", "0003", "0004"},
			Turn: &deviant.Turn{
				Id:    "turn_0000",
				Phase: deviant.TurnPhaseNames_PHASE_POINT,
			},
		},
	}

	var marshalOptions = protojson.MarshalOptions{
		AllowPartial:    true,
		EmitUnpopulated: true,
	}

	result, _ := protojson.MarshalOptions(marshalOptions).Marshal(test)
	redisClient := NewCacheClient()
	err := redisClient.Set("encounter_0000", string(result), 0).Err()
	if err != nil {
		panic(err)
	}

	return test
}
