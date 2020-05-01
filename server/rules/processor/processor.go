package processor

import (
	"github.com/recluse-games/deviant-instance-shard/server/rules/deck"
	"github.com/recluse-games/deviant-instance-shard/server/rules/hand"
	"github.com/recluse-games/deviant-instance-shard/server/rules/turn"
	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

// Process Processes all rules to determine validity.
func Process(turnPhaseName deviant.TurnPhaseNames, entity *deviant.Entity, entityAction deviant.EntityActionNames) bool {
	entityActionRules := map[deviant.EntityActionNames][]func(*deviant.Entity) bool{
		deviant.EntityActionNames_PLAY:    {},
		deviant.EntityActionNames_DRAW:    {hand.ValidateSize, deck.ValidateDraw},
		deviant.EntityActionNames_DISCARD: {},
	}

	turnPhaseRules := map[deviant.TurnPhaseNames][]func(deviant.EntityActionNames, deviant.TurnPhaseNames) bool{
		deviant.TurnPhaseNames_PHASE_POINT:   {turn.ValidateEntityAction},
		deviant.TurnPhaseNames_PHASE_EFFECT:  {turn.ValidateEntityAction},
		deviant.TurnPhaseNames_PHASE_ACTION:  {turn.ValidateEntityAction},
		deviant.TurnPhaseNames_PHASE_DISCARD: {turn.ValidateEntityAction},
		deviant.TurnPhaseNames_PHASE_END:     {turn.ValidateEntityAction},
	}

	if val, ok := entityActionRules[entityAction]; ok {
		if ok {
			for _, entityRuleFunction := range val {
				if entityRuleFunction(entity) == false {
					return false
				}
			}
		} else {
			return false
		}
	}

	if val, ok := turnPhaseRules[turnPhaseName]; ok {
		if ok {
			for _, turnRuleFunction := range val {
				if turnRuleFunction(entityAction, turnPhaseName) == false {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}
