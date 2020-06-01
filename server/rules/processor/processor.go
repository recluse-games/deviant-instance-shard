package processor

import (
	"fmt"

	"github.com/golang/glog"
	deck "github.com/recluse-games/deviant-instance-shard/server/rules/deck"
	hand "github.com/recluse-games/deviant-instance-shard/server/rules/hand"
	move "github.com/recluse-games/deviant-instance-shard/server/rules/move"
	play "github.com/recluse-games/deviant-instance-shard/server/rules/play"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Process Processes all rules to determine validity.
func Process(encounter *deviant.Encounter, entityActionName deviant.EntityActionNames, entityMoveAction *deviant.EntityMoveAction, entityPlayAction *deviant.EntityPlayAction) bool {
	entityActionRules := map[deviant.EntityActionNames][]interface{}{
		deviant.EntityActionNames_PLAY:         {play.ValidateApCost},
		deviant.EntityActionNames_MOVE:         {move.ValidateApCost, move.ValidateNewLocationEmpty, move.ValidateNewLocationSide},
		deviant.EntityActionNames_DISCARD:      {},
		deviant.EntityActionNames_CHANGE_PHASE: {},
	}

	turnPhaseRules := map[deviant.TurnPhaseNames][]interface{}{
		deviant.TurnPhaseNames_PHASE_POINT:   {},
		deviant.TurnPhaseNames_PHASE_EFFECT:  {},
		deviant.TurnPhaseNames_PHASE_DRAW:    {hand.ValidateSize, deck.ValidateDraw},
		deviant.TurnPhaseNames_PHASE_ACTION:  {},
		deviant.TurnPhaseNames_PHASE_DISCARD: {},
		deviant.TurnPhaseNames_PHASE_END:     {},
	}

	validEntityActionsPerTurn := map[deviant.TurnPhaseNames]map[deviant.EntityActionNames]bool{
		deviant.TurnPhaseNames_PHASE_POINT: {
			deviant.EntityActionNames_NOTHING: true,
		},
		deviant.TurnPhaseNames_PHASE_EFFECT: {
			deviant.EntityActionNames_NOTHING: true,
		},
		deviant.TurnPhaseNames_PHASE_DRAW: {
			deviant.EntityActionNames_NOTHING: true,
		},
		deviant.TurnPhaseNames_PHASE_ACTION: {
			deviant.EntityActionNames_NOTHING:      true,
			deviant.EntityActionNames_PLAY:         true,
			deviant.EntityActionNames_MOVE:         true,
			deviant.EntityActionNames_CHANGE_PHASE: true,
		},
		deviant.TurnPhaseNames_PHASE_DISCARD: {
			deviant.EntityActionNames_NOTHING: true,
			deviant.EntityActionNames_DISCARD: true,
		},
		deviant.TurnPhaseNames_PHASE_END: {
			deviant.EntityActionNames_NOTHING: true,
		},
	}

	// Early return for any incoming actions that don't match turns.
	if _, ok := validEntityActionsPerTurn[encounter.Turn.Phase][entityActionName]; ok {
		if !ok {
			return false
		}
	}

	// Validate all entity action rules are sucessfully passing.
	if val, ok := entityActionRules[entityActionName]; ok {
		if ok {
			for _, entityActionFunction := range val {
				switch entityActionName {
				case deviant.EntityActionNames_MOVE:
					if entityActionFunction.(func(*deviant.Entity, *deviant.EntityMoveAction, *deviant.Encounter) bool)(encounter.ActiveEntity, entityMoveAction, encounter) == false {
						return false
					}
				case deviant.EntityActionNames_PLAY:
					if entityActionFunction.(func(*deviant.Entity, *deviant.EntityPlayAction, *deviant.Encounter) bool)(encounter.ActiveEntity, entityPlayAction, encounter) == false {
						return false
					}
				case deviant.EntityActionNames_CHANGE_PHASE:
					if entityActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
						return false
					}
				default:
					message := fmt.Sprintf("No rules implemented implemented for EntityActionName: %s", entityActionName.String())
					glog.Error(message)
				}
			}
		} else {
			message := fmt.Sprintf("Invalid EntityActionName: %s", entityActionName)
			glog.Error(message)
			return false
		}
	}

	// Validate all turn phase rules are sucessfully passing.
	if val, ok := turnPhaseRules[encounter.Turn.Phase]; ok {
		if ok {
			for _, turnRuleFunction := range val {
				switch encounter.Turn.Phase {
				case deviant.TurnPhaseNames_PHASE_POINT:
				case deviant.TurnPhaseNames_PHASE_EFFECT:
				case deviant.TurnPhaseNames_PHASE_DRAW:
					if turnRuleFunction.(func(*deviant.Encounter) bool)(encounter) == false {
						return false
					}
				case deviant.TurnPhaseNames_PHASE_ACTION:
				case deviant.TurnPhaseNames_PHASE_DISCARD:
				case deviant.TurnPhaseNames_PHASE_END:
				default:
					message := fmt.Sprintf("No rules implemented for TurnPhaseName: %s", entityActionName.String())
					glog.Error(message)
				}
			}
		} else {
			message := fmt.Sprintf("Invalid TurnPhaseName: %s", encounter.Turn.Phase)
			glog.Error(message)
			return false
		}
	}

	return true
}
