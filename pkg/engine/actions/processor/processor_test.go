package processor

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestProcess(t *testing.T) {
	turn := &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	entity := &deviant.Entity{
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
	if Process(encounter, deviant.EntityActionNames_NOTHING, nil, nil, nil) != true {
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
	if Process(encounter, deviant.EntityActionNames_NOTHING, nil, nil, nil) != true {
		t.Fail()
	}

	if encounter.ActiveEntity.Ap != 5 {
		t.Logf("%v", encounter.ActiveEntity)
		t.Fail()
	}
}
