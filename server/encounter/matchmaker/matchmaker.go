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

		card = &deviant.Card{
			Id:          "attack_slash_0000",
			BackId:      "back_0000",
			InstanceId:  uuid.New().String(),
			Cost:        2,
			Damage:      3,
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
								Distance:  2,
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
				OverlayTiles: []*deviant.Tile{},
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
									Id:         "0002",
									Hp:         10,
									Ap:         4,
									Alignment:  deviant.Alignment_UNFRIENDLY,
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
			ActiveEntityOrder: []string{"0001", "0002"},
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
