package actions

import (
	"fmt"

	"github.com/golang/glog"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Process Processes all actions.
func Process(encounter *deviant.Encounter, entityActionName deviant.EntityActionNames, entityMoveAction *deviant.EntityMoveAction, entityPlayAction *deviant.EntityPlayAction, entityRotateAction *deviant.EntityRotateAction) bool {
	entityActions := map[deviant.EntityActionNames][]interface{}{
		deviant.EntityActionNames_PLAY:         {Play},
		deviant.EntityActionNames_MOVE:         {Move},
		deviant.EntityActionNames_ROTATE:       {Rotate},
		deviant.EntityActionNames_DRAW:         {},
		deviant.EntityActionNames_DISCARD:      {},
		deviant.EntityActionNames_NOTHING:      {},
		deviant.EntityActionNames_CHANGE_PHASE: {ChangePhase},
	}

	turnActions := map[deviant.TurnPhaseNames][]interface{}{
		deviant.TurnPhaseNames_PHASE_POINT:   {GrantAp, ChangePhase},
		deviant.TurnPhaseNames_PHASE_EFFECT:  {ChangePhase},
		deviant.TurnPhaseNames_PHASE_DRAW:    {DrawCard, ChangePhase},
		deviant.TurnPhaseNames_PHASE_ACTION:  {},
		deviant.TurnPhaseNames_PHASE_DISCARD: {},
		deviant.TurnPhaseNames_PHASE_END:     {UpdateActiveEntity, ChangePhase},
	}

	encounterActions := []func(*deviant.Encounter) bool{
		ProcessWinConditions,
	}

	if val, ok := entityActions[entityActionName]; ok {
		if ok {
			for _, entityActionFunction := range val {
				switch entityActionName {
				case deviant.EntityActionNames_PLAY:
					if entityActionFunction.(func(*deviant.Encounter, *deviant.EntityPlayAction) bool)(encounter, entityPlayAction) == false {
						return false
					}
				case deviant.EntityActionNames_MOVE:
					if entityActionFunction.(func(*deviant.Encounter, *deviant.EntityMoveAction) bool)(encounter, entityMoveAction) == false {
						return false
					}
				case deviant.EntityActionNames_CHANGE_PHASE:
					if entityActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
						return false
					}
				case deviant.EntityActionNames_ROTATE:
					if entityActionFunction.(func(*deviant.Encounter, *deviant.EntityRotateAction) bool)(encounter, entityRotateAction) == false {
						return false
					}
				default:
					message := fmt.Sprintf("No actions implemented for EntityActionName: %s", entityActionName.String())
					glog.Error(message)
				}
			}
		} else {
			message := fmt.Sprintf("Invalid EntityActionName: %s", entityActionName)
			glog.Error(message)
			return false
		}
	}

ProcessTurnActions:
	// Wrapped forever loop allows for server-side phase skipping/processing.
	for {
		if val, ok := turnActions[encounter.Turn.Phase]; ok {
			if ok {
				var nextTurnPhaseName deviant.TurnPhaseNames

				// If we're on the end phase we need to loop back to PHASE_POINT in the TurnPhaseActionNames.
				if encounter.Turn.Phase < 5 {
					nextTurnPhaseName = encounter.Turn.Phase + 1
				} else {
					nextTurnPhaseName = deviant.TurnPhaseNames_PHASE_POINT
				}

				for _, turnActionFunction := range val {
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
						if turnActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
							return false
						}
					case deviant.TurnPhaseNames_PHASE_END:
						if turnActionFunction.(func(*deviant.Encounter) bool)(encounter) == false {
							return false
						}
					default:
						message := fmt.Sprintf("No actions implemented for EntityActionName: %s", entityActionName.String())
						glog.Error(message)
					}

					// If one of our previously executed action processors incremented the phase we should loop.
					if encounter.Turn.Phase == nextTurnPhaseName {
						continue ProcessTurnActions
					}
				}

				// If we're not above the maximum hand size we should skip processing.
				if encounter.Turn.Phase == deviant.TurnPhaseNames_PHASE_DISCARD && len(encounter.ActiveEntity.Hand.Cards) < 6 {
					ChangePhase(encounter)
					continue ProcessTurnActions
				}

				break ProcessTurnActions
			} else {
				message := fmt.Sprintf("Invalid EntityActionName: %s", entityActionName)
				glog.Error(message)
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
