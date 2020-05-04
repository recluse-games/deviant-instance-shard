package processor

import (
	"log"

	"github.com/recluse-games/deviant-instance-shard/server/rules/deck"
	"github.com/recluse-games/deviant-instance-shard/server/rules/hand"
	"github.com/recluse-games/deviant-instance-shard/server/rules/move"
	"github.com/recluse-games/deviant-instance-shard/server/rules/turn"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Process Processes all rules to determine validity.
func Process(encounter *deviant.Encounter, entityAction deviant.EntityActionNames, entityMoveAction *deviant.EntityMoveAction) bool {
	entityActionRules := map[deviant.EntityActionNames][]interface{}{
		deviant.EntityActionNames_PLAY:         {},
		deviant.EntityActionNames_MOVE:         {move.ValidateApCost},
		deviant.EntityActionNames_DISCARD:      {},
		deviant.EntityActionNames_CHANGE_PHASE: {},
	}

	turnPhaseRules := map[deviant.TurnPhaseNames][]func(*deviant.Encounter, deviant.EntityActionNames, deviant.TurnPhaseNames) bool{
		deviant.TurnPhaseNames_PHASE_POINT:   {turn.ValidateEntityAction},
		deviant.TurnPhaseNames_PHASE_EFFECT:  {turn.ValidateEntityAction},
		deviant.TurnPhaseNames_PHASE_DRAW:    {turn.ValidateEntityAction, hand.ValidateSize, deck.ValidateDraw},
		deviant.TurnPhaseNames_PHASE_ACTION:  {turn.ValidateEntityAction},
		deviant.TurnPhaseNames_PHASE_DISCARD: {turn.ValidateEntityAction},
		deviant.TurnPhaseNames_PHASE_END:     {turn.ValidateEntityAction},
	}

	if val, ok := entityActionRules[entityAction]; ok {
		if ok {
			for actionName, entityActionFunction := range val {
				switch entityAction {
				case deviant.EntityActionNames_MOVE:
					if entityActionFunction.(func(*deviant.Entity, *deviant.EntityMoveAction) bool)(encounter.ActiveEntity, entityMoveAction) == false {
						return false
					}
				case deviant.EntityActionNames_CHANGE_PHASE:
					if entityActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
						return false
					}
				default:
					log.Output(1, deviant.EntityActionNames_name[int32(actionName)])
				}
			}
		} else {
			return false
		}
	}

	if val, ok := turnPhaseRules[encounter.Turn.Phase]; ok {
		if ok {
			for _, turnRuleFunction := range val {
				if turnRuleFunction(encounter, entityAction, encounter.Turn.Phase) == false {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}
