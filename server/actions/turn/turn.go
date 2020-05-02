package turn

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// GrantAp grants some entity a default AP value of 5
func GrantAp(encounter *deviant.Encounter) bool {
	startingApValue := int32(5)
	encounter.ActiveEntity.Ap = startingApValue

	return true
}

// ChangePhase updates the current turn phase.
func ChangePhase(encounter *deviant.Encounter) bool {
	turnOrder := []deviant.TurnPhaseNames{deviant.TurnPhaseNames_PHASE_POINT, deviant.TurnPhaseNames_PHASE_DRAW, deviant.TurnPhaseNames_PHASE_EFFECT, deviant.TurnPhaseNames_PHASE_ACTION, deviant.TurnPhaseNames_PHASE_DISCARD, deviant.TurnPhaseNames_PHASE_END}
	for i, v := range turnOrder {
		if v == encounter.Turn.Phase {
			encounter.Turn.Phase = turnOrder[i+1]
			break
		}
	}

	return true
}

// UpdateActiveEntity Updates the active entity to the next entity in the active entity order.
func UpdateActiveEntity(encounter *deviant.Encounter) bool {
	var newActiveEntity *deviant.Entity

	for index, entityID := range encounter.ActiveEntityOrder {
		if entityID == encounter.ActiveEntity.Id {
			newActiveEntity.Id = encounter.ActiveEntityOrder[index+1]
		}
	}

	// Apply all state changes to entity in encounter as well as the activeEntity
	for outerIndex, outerValue := range encounter.Board.Entities.Entities {
		for innerIndex, innerValue := range outerValue.Entities {
			if innerValue != nil {
				if innerValue.Id == newActiveEntity.Id {
					encounter.ActiveEntity = encounter.Board.Entities.Entities[outerIndex].Entities[innerIndex]
				}
			}
		}
	}

	return true
}
