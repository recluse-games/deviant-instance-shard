package actions

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestGrantAp(t *testing.T) {
	entityWithNoAp := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Ap:    0,
			MaxAp: 5,
		},
	}

	if GrantAp(entityWithNoAp, nil) != true {
		t.Fail()
	}

	if entityWithNoAp.ActiveEntity.Ap != 5 {
		t.Fail()
	}
}

func TestChangePhase(t *testing.T) {
	// Test Non End Phase Incrementation
	mockTurn := &deviant.Turn{
		Id:    "0000",
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	mockEncounter := &deviant.Encounter{
		Turn: mockTurn,
	}

	ChangePhase(mockEncounter, nil)

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
		Turn: mockTurn,
	}

	ChangePhase(mockEncounter, nil)

	if mockEncounter.Turn.Phase != deviant.TurnPhaseNames_PHASE_POINT {
		t.Logf("Failed to increment the turn correctly from end")
		t.Fail()
	}

}

func TestUpdateActiveEntity(t *testing.T) {
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
							&deviant.Entity{
								Hp: 5,
								Id: "0001",
							},
							&deviant.Entity{
								Hp: 5,
								Id: "0002",
							},
						},
					},
				},
			},
		},
	}

	result := UpdateActiveEntity(mockEncounter, nil)

	if result == false {
		t.Logf("Failed to update active entity")
		t.Fail()
	}
}
