package actions

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap/zaptest"
)

func TestProcess(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	turn := &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	entity := &deviant.Entity{
		Hp:    10,
		Ap:    0,
		MaxAp: 5,
		Deck:  enginetest.GenerateDeckLiteral(5),
		Hand:  enginetest.GenerateHandLiteral(5),
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
	if Process(encounter, deviant.EntityActionNames_NOTHING, nil, nil, nil, logger.Sugar()) != true {
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
	if Process(encounter, deviant.EntityActionNames_NOTHING, nil, nil, nil, logger.Sugar()) != true {
		t.Fail()
	}

	if encounter.ActiveEntity.Ap != 5 {
		t.Logf("%v", encounter.ActiveEntity)
		t.Fail()
	}

	turn = &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	entity = &deviant.Entity{
		Id:    "0001",
		Hp:    1,
		Ap:    0,
		MaxAp: 5,
		Deck:  enginetest.GenerateDeckLiteral(0),
		Hand:  enginetest.GenerateHandLiteral(0),
	}

	entity2 := &deviant.Entity{
		Id:    "0002",
		Hp:    1,
		Ap:    0,
		MaxAp: 5,
		Deck:  enginetest.GenerateDeckLiteral(1),
		Hand:  enginetest.GenerateHandLiteral(1),
	}

	encounter = &deviant.Encounter{
		ActiveEntity:      entity,
		Turn:              turn,
		ActiveEntityOrder: []string{"0001", "0002"},
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
							entity2,
						},
					},
				},
			},
		},
	}

	// Test GrantAP Turn Phase Action
	if Process(encounter, deviant.EntityActionNames_NOTHING, nil, nil, nil, logger.Sugar()) != true {
		t.Fail()
	}

	if encounter.ActiveEntityOrder[0] != "0002" {
		t.Logf("%v", encounter.ActiveEntity)
		t.Fail()
	}

	if encounter.ActiveEntity.Id != "0002" {
		t.Logf("%v", encounter.ActiveEntity)
		t.Fail()
	}
}
