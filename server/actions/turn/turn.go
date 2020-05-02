package turn

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// GrantAp grants some entity a default AP value of 5
func GrantAp(entity *deviant.Entity) bool {
	startingApValue := int32(5)
	entity.Ap = startingApValue

	return true
}

// ChangePhase updates the current turn phase.
func ChangePhase(entity *deviant.Entity, turn *deviant.Turn) bool {
	turnOrder := []deviant.TurnPhaseNames{deviant.TurnPhaseNames_PHASE_POINT, deviant.TurnPhaseNames_PHASE_DRAW, deviant.TurnPhaseNames_PHASE_EFFECT, deviant.TurnPhaseNames_PHASE_ACTION, deviant.TurnPhaseNames_PHASE_DISCARD, deviant.TurnPhaseNames_PHASE_END}
	for i, v := range turnOrder {
		if v == turn.Phase {
			turn.Phase = turnOrder[i+1]
			break
		}
	}

	return true
}
