package turn

import deviant "github.com/recluse-games/deviant-protobuf/genproto/go"

// ValidateEntityAction Determines if the action an entity is trying to perform is valid for the current turnPhase
func ValidateEntityAction(encounter *deviant.Encounter, actionName deviant.EntityActionNames, turnPhaseName deviant.TurnPhaseNames) bool {
	validEntityTurnPhasePointActions := map[deviant.EntityActionNames]bool{
		deviant.EntityActionNames_NOTHING:      true,
		deviant.EntityActionNames_CHANGE_PHASE: true,
	}
	validEntityTurnPhaseEffectActions := map[deviant.EntityActionNames]bool{
		deviant.EntityActionNames_NOTHING:      true,
		deviant.EntityActionNames_CHANGE_PHASE: true,
	}
	validEntityTurnPhaseDrawActions := map[deviant.EntityActionNames]bool{
		deviant.EntityActionNames_NOTHING:      true,
		deviant.EntityActionNames_CHANGE_PHASE: true,
	}
	validEntityTurnPhaseActionActions := map[deviant.EntityActionNames]bool{
		deviant.EntityActionNames_NOTHING:      true,
		deviant.EntityActionNames_PLAY:         true,
		deviant.EntityActionNames_CHANGE_PHASE: true,
	}
	validEntityTurnPhaseDiscardActions := map[deviant.EntityActionNames]bool{
		deviant.EntityActionNames_NOTHING:      true,
		deviant.EntityActionNames_DISCARD:      true,
		deviant.EntityActionNames_CHANGE_PHASE: true,
	}
	validEntityTurnPhaseEndActions := map[deviant.EntityActionNames]bool{
		deviant.EntityActionNames_NOTHING:      true,
		deviant.EntityActionNames_CHANGE_PHASE: true,
	}

	validEntityActions := map[deviant.TurnPhaseNames]map[deviant.EntityActionNames]bool{
		deviant.TurnPhaseNames_PHASE_POINT:   validEntityTurnPhasePointActions,
		deviant.TurnPhaseNames_PHASE_EFFECT:  validEntityTurnPhaseEffectActions,
		deviant.TurnPhaseNames_PHASE_DRAW:    validEntityTurnPhaseDrawActions,
		deviant.TurnPhaseNames_PHASE_ACTION:  validEntityTurnPhaseActionActions,
		deviant.TurnPhaseNames_PHASE_DISCARD: validEntityTurnPhaseDiscardActions,
		deviant.TurnPhaseNames_PHASE_END:     validEntityTurnPhaseEndActions,
	}

	if _, ok := validEntityActions[turnPhaseName][actionName]; ok {
		if ok {
			return true
		}
	}

	return false
}
