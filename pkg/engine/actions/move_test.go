package actions

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap/zaptest"
)

func TestMove(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	mockEntity := &deviant.Entity{
		Id:   "0001",
		Ap:   2,
		Deck: enginetest.GenerateDeckLiteral(1),
		Hand: enginetest.GenerateHandLiteral(1),
	}

	mockEncounter := &deviant.Encounter{
		ActiveEntity: mockEntity,
		Board: &deviant.Board{
			Entities: &deviant.Entities{
				Entities: []*deviant.EntitiesRow{
					{
						Entities: []*deviant.Entity{
							mockEntity,
							{},
						},
					},
					{
						Entities: []*deviant.Entity{
							{},
							{},
						},
					},
				},
			},
		},
	}

	mockMoveAction := &deviant.EntityMoveAction{
		StartXPosition: 0,
		StartYPosition: 0,
		FinalXPosition: 1,
		FinalYPosition: 1,
	}

	if Move(mockEncounter, mockMoveAction, logger.Sugar()) != true {
		t.Fail()
	}

	if mockEncounter.ActiveEntity.Ap != 0 {
		t.Logf("Failed to consume entity AP properly")
		t.Fail()
	}

	if mockEncounter.Board.Entities.Entities[1].Entities[1].Id != "0001" {
		t.Logf("Failed to move entity to correct location")
		t.Fail()
	}
}
