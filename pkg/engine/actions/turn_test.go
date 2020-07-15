package actions

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap/zaptest"
)

func TestGrantAp(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entityWithNoAp := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Id:    "0001",
			Ap:    0,
			MaxAp: 5,
		},
	}

	if GrantAp(entityWithNoAp, logger.Sugar()) != true {
		t.Fail()
	}

	if entityWithNoAp.ActiveEntity.Ap != 5 {
		t.Fail()
	}
}

func TestChangePhase(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	// Test Non End Phase Incrementation
	mockTurn := &deviant.Turn{
		Id:    "0000",
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	mockEncounter := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Id: "0001",
		},
		Turn: mockTurn,
	}

	ChangePhase(mockEncounter, logger.Sugar())

	if mockEncounter.Turn.Phase != deviant.TurnPhaseNames_PHASE_ACTION {
		t.Logf("Failed to increment the turn phase")
		t.Fail()
	}

	// Test End Phase Incrementation
	mockTurn = &deviant.Turn{
		Id:    "0000",
		Phase: deviant.TurnPhaseNames_PHASE_END,
	}

	mockEncounter = &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Id: "0001",
		},
		Turn: mockTurn,
	}

	ChangePhase(mockEncounter, logger.Sugar())

	if mockEncounter.Turn.Phase != deviant.TurnPhaseNames_PHASE_POINT {
		t.Logf("Failed to increment the turn correctly from end")
		t.Fail()
	}

}

func TestUpdateActiveEntity(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	mockTurn := &deviant.Turn{
		Id:    "0000",
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	mockEncounter := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Hp: 5,
			Id: "0001",
		},
		ActiveEntityOrder: []string{"0001", "0002", "0003"},
		Turn:              mockTurn,
		Board: &deviant.Board{
			Entities: &deviant.Entities{
				Entities: []*deviant.EntitiesRow{
					{
						Entities: []*deviant.Entity{
							{
								Hp: 5,
								Id: "0001",
							},
							{
								Hp: 5,
								Id: "0002",
							},
						},
					},
				},
			},
		},
	}

	result := UpdateActiveEntity(mockEncounter, logger.Sugar())

	if result == false {
		t.Logf("Failed to update active entity")
		t.Fail()
	}
}
