package processor

import (
	deckActions "github.com/recluse-games/deviant-instance-shard/server/actions/deck"
	encounterActions "github.com/recluse-games/deviant-instance-shard/server/actions/encounter"
	"github.com/recluse-games/deviant-instance-shard/server/actions/turn"
	turnActions "github.com/recluse-games/deviant-instance-shard/server/actions/turn"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Process Processes all actions.
func Process(encounter *deviant.Encounter, entityActionName deviant.EntityActionNames) bool {
	entityActions := map[deviant.EntityActionNames][]func(*deviant.Encounter) bool{
		deviant.EntityActionNames_PLAY:         {},
		deviant.EntityActionNames_DISCARD:      {},
		deviant.EntityActionNames_NOTHING:      {},
		deviant.EntityActionNames_CHANGE_PHASE: {turnActions.ChangePhase},
	}

	turnActions := map[deviant.TurnPhaseNames][]func(*deviant.Encounter) bool{
		deviant.TurnPhaseNames_PHASE_POINT:   {turnActions.GrantAp},
		deviant.TurnPhaseNames_PHASE_DRAW:    {deckActions.DrawCard},
		deviant.TurnPhaseNames_PHASE_EFFECT:  {},
		deviant.TurnPhaseNames_PHASE_ACTION:  {},
		deviant.TurnPhaseNames_PHASE_DISCARD: {},
		deviant.TurnPhaseNames_PHASE_END:     {turn.UpdateActiveEntity},
	}

	encounterActions := []func(*deviant.Encounter) bool{
		encounterActions.ProcessWinConditions,
	}

	if val, ok := entityActions[entityActionName]; ok {
		if ok {
			for _, entityActionFunction := range val {
				if entityActionFunction(encounter) == false {
					return false
				}
			}
		} else {
			return false
		}
	}

	if val, ok := turnActions[encounter.Turn.Phase]; ok {
		if ok {
			for _, turnActionFunction := range val {
				if turnActionFunction(encounter) == false {
					return false
				}
			}
		} else {
			return false
		}
	}

	for _, encounterActionFunction := range encounterActions {
		if encounterActionFunction(encounter) == false {
			return false
		}
	}

	return true
}
