package processor

import (
	"github.com/recluse-games/deviant-instance-shard/server/actions/deck"
	"github.com/recluse-games/deviant-instance-shard/server/actions/turn"
	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

// Process Processes all actions.
func Process(selectedTurn *deviant.Turn, selectedEntity *deviant.Entity, entityActionName deviant.EntityActionNames) bool {
	entityActions := map[deviant.EntityActionNames][]func(*deviant.Entity) bool{
		deviant.EntityActionNames_PLAY:    {},
		deviant.EntityActionNames_DISCARD: {},
		deviant.EntityActionNames_NOTHING: {},
	}

	turnActions := map[deviant.TurnPhaseNames][]func(*deviant.Entity) bool{
		deviant.TurnPhaseNames_PHASE_POINT:   {turn.GrantAp},
		deviant.TurnPhaseNames_PHASE_DRAW:    {deck.DrawCard},
		deviant.TurnPhaseNames_PHASE_EFFECT:  {},
		deviant.TurnPhaseNames_PHASE_ACTION:  {},
		deviant.TurnPhaseNames_PHASE_DISCARD: {},
		deviant.TurnPhaseNames_PHASE_END:     {},
	}

	if val, ok := entityActions[entityActionName]; ok {
		if ok {
			for _, entityActionFunction := range val {
				if entityActionFunction(selectedEntity) == false {
					return false
				}
			}
		} else {
			return false
		}
	}

	if val, ok := turnActions[selectedTurn.Phase]; ok {
		if ok {
			for _, turnActionFunction := range val {
				if turnActionFunction(selectedEntity) == false {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}
