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

	turnActions := map[deviant.TurnPhaseNames][]interface{}{
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

ProcessTurnActions:
	// Wrapped forever loop allows for server-side phase skipping/processing.
	for {
		if val, ok := turnActions[encounter.Turn.Phase]; ok {
			if ok {
				for _, turnActionFunction := range val {
					nextTurnPhaseName := encounter.Turn.Phase + 1

					switch encounter.Turn.Phase {
					case deviant.TurnPhaseNames_PHASE_POINT:
						if turnActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
							return false
						}
					case deviant.TurnPhaseNames_PHASE_EFFECT:
						if turnActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
							return false
						}
					case deviant.TurnPhaseNames_PHASE_DRAW:
						if turnActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
							return false
						}
					case deviant.TurnPhaseNames_PHASE_ACTION:
						if turnActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
							return false
						}
					case deviant.TurnPhaseNames_PHASE_DISCARD:
						// If we're not above the maximum hand size we should skip processing.
						if len(encounter.ActiveEntity.Hand.Cards) < 6 {
							turn.ChangePhase(encounter)
							continue ProcessTurnActions
						}
					case deviant.TurnPhaseNames_PHASE_END:
						if turnActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
							return false
						}
					default:
						log.Fatal("Request TurnPhaseName is not implemented.")
					}

					// If one of our previously executed action processors incremented the phase we should loop.
					if encounter.Turn.Phase == nextTurnPhaseName {
						continue ProcessTurnActions
					}
				}

				break ProcessTurnActions
			} else {
				return false
			}
		}
	}

	// Verify all actions were completed succesfully.
	for _, encounterActionFunction := range encounterActions {
		if encounterActionFunction(encounter) == false {
			return false
		}
	}

	return true
}
