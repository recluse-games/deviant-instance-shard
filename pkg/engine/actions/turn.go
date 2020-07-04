package actions

import (
	"fmt"

	"github.com/golang/glog"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// GrantAp grants some entity a default AP value of 5
func GrantAp(encounter *deviant.Encounter) bool {
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.MaxAp

	message := fmt.Sprintf("Action: Grant AP: %s", encounter.ActiveEntity.Id)
	glog.Info(message)

	return true
}

// ChangePhase updates the current turn phase.
func ChangePhase(encounter *deviant.Encounter) bool {
	turnOrder := []deviant.TurnPhaseNames{deviant.TurnPhaseNames_PHASE_POINT, deviant.TurnPhaseNames_PHASE_EFFECT, deviant.TurnPhaseNames_PHASE_DRAW, deviant.TurnPhaseNames_PHASE_ACTION, deviant.TurnPhaseNames_PHASE_DISCARD, deviant.TurnPhaseNames_PHASE_END}

	for i, v := range turnOrder {
		if v == encounter.Turn.Phase {
			if encounter.Turn.Phase == deviant.TurnPhaseNames_PHASE_END {
				encounter.Turn.Phase = turnOrder[0]
				break
			}

			encounter.Turn.Phase = turnOrder[i+1]

			break
		}
	}

	message := fmt.Sprintf("Action: Changed Phase: %s", encounter.Turn.Phase)
	glog.Info(message)

	return true
}

// UpdateActiveEntity Updates the active entity to the next entity in the active entity order.
func UpdateActiveEntity(encounter *deviant.Encounter) bool {
	var newActiveEntityID string

	for index, entityID := range encounter.ActiveEntityOrder {
		if entityID == encounter.ActiveEntity.Id {
			if len(encounter.ActiveEntityOrder) != index+1 {
				newActiveEntityID = encounter.ActiveEntityOrder[index+1]
			} else {
				newActiveEntityID = encounter.ActiveEntityOrder[0]
			}
		}
	}

	for _, entitiesRow := range encounter.Board.Entities.Entities {
		for _, entity := range entitiesRow.Entities {
			if entity.Id == newActiveEntityID {
				encounter.ActiveEntity = entity
			}
		}
	}

	message := fmt.Sprintf("Action: Updated Active Entity: %s", encounter.ActiveEntity.Id)
	glog.Info(message)

	return true
}
