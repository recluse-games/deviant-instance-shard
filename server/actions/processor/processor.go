package processor

import (
	"log"

	deckActions "github.com/recluse-games/deviant-instance-shard/server/actions/deck"
	encounterActions "github.com/recluse-games/deviant-instance-shard/server/actions/encounter"
	moveActions "github.com/recluse-games/deviant-instance-shard/server/actions/move"
	"github.com/recluse-games/deviant-instance-shard/server/actions/turn"
	turnActions "github.com/recluse-games/deviant-instance-shard/server/actions/turn"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Process Processes all actions.
func Process(encounter *deviant.Encounter, entityActionName deviant.EntityActionNames, entityMoveAction *deviant.EntityMoveAction) bool {
	entityActions := map[deviant.EntityActionNames][]interface{}{
		deviant.EntityActionNames_PLAY:         {},
		deviant.EntityActionNames_MOVE:         {moveActions.Move},
		deviant.EntityActionNames_DRAW:         {},
		deviant.EntityActionNames_DISCARD:      {},
		deviant.EntityActionNames_NOTHING:      {},
		deviant.EntityActionNames_CHANGE_PHASE: {turn.ChangePhase},
	}

	turnActions := map[deviant.TurnPhaseNames][]func(*deviant.Encounter) bool{
		deviant.TurnPhaseNames_PHASE_POINT:   {turnActions.GrantAp, turn.ChangePhase},
		deviant.TurnPhaseNames_PHASE_EFFECT:  {turn.ChangePhase},
		deviant.TurnPhaseNames_PHASE_DRAW:    {deckActions.DrawCard, turn.ChangePhase},
		deviant.TurnPhaseNames_PHASE_ACTION:  {},
		deviant.TurnPhaseNames_PHASE_DISCARD: {},
		deviant.TurnPhaseNames_PHASE_END:     {turn.UpdateActiveEntity, turn.ChangePhase},
	}

	encounterActions := []func(*deviant.Encounter) bool{
		encounterActions.ProcessWinConditions,
	}

	if val, ok := entityActions[entityActionName]; ok {
		if ok {
			for actionName, entityActionFunction := range val {
				switch entityActionName {
				case deviant.EntityActionNames_MOVE:
					if entityActionFunction.(func(*deviant.Encounter, *deviant.EntityMoveAction) bool)(encounter, entityMoveAction) == false {
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

ProcessTurns:
	// Wrapped loop always turnphases to change from other actions to skip certain phases.
	for {
		if val, ok := turnActions[encounter.Turn.Phase]; ok {
			if ok {
				for _, turnActionFunction := range val {
					var nextTurnName = encounter.Turn.Phase + 1

					if turnActionFunction(encounter) == false {
						return false
					}

					if encounter.Turn.Phase == nextTurnName {
						continue ProcessTurns
					}

					// Edge Case skip discard phase if not neccisary
					if len(encounter.ActiveEntity.Hand.Cards) < 6 {
						turn.ChangePhase(encounter)
						continue ProcessTurns
					}
				}

				break ProcessTurns
			} else {
				return false
			}
		}

		break ProcessTurns
	}

	for _, encounterActionFunction := range encounterActions {
		if encounterActionFunction(encounter) == false {
			return false
		}
	}

	return true
}
