package actions

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// GrantAp grants some entity a default AP value of 5
func GrantAp(encounter *deviant.Encounter, logger *zap.Logger) bool {
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.MaxAp

	if logger != nil {
		logger.Debug("Entity Granted AP",
			zap.String("actionID", "GrantAP"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return true
}

// ChangePhase updates the current turn phase.
func ChangePhase(encounter *deviant.Encounter, logger *zap.Logger) bool {
	turnOrder := []deviant.TurnPhaseNames{deviant.TurnPhaseNames_PHASE_POINT, deviant.TurnPhaseNames_PHASE_EFFECT, deviant.TurnPhaseNames_PHASE_DRAW, deviant.TurnPhaseNames_PHASE_ACTION, deviant.TurnPhaseNames_PHASE_DISCARD, deviant.TurnPhaseNames_PHASE_END}

	for i, v := range turnOrder {
		if v == encounter.Turn.Phase {
			if encounter.Turn.Phase == deviant.TurnPhaseNames_PHASE_END {
				encounter.Turn.Phase = turnOrder[0]
				break
			}

			encounter.Turn.Phase = turnOrder[i+1]

			break
		}
	}

	if logger != nil {
		logger.Debug("Phase Changed",
			zap.String("actionID", "ChangePhase"),
			zap.String("entityID", encounter.ActiveEntity.Id),
			zap.String("phaseID", encounter.Turn.Phase.String()),
		)
	}

	return true
}

// UpdateActiveEntity Updates the active entity to the next entity in the active entity order.
func UpdateActiveEntity(encounter *deviant.Encounter, logger *zap.Logger) bool {
	var newActiveEntityID string

	newActiveEntityIndex, _ := findSliceIndex(encounter.ActiveEntityOrder, encounter.ActiveEntity.Id)

	if len(encounter.ActiveEntityOrder) > newActiveEntityIndex+1 {
		newActiveEntityID = encounter.ActiveEntityOrder[newActiveEntityIndex+1]
	} else if len(encounter.ActiveEntityOrder) == 1 {
		newActiveEntityID = encounter.ActiveEntityOrder[0]
	}

	if encounter.ActiveEntity.Hp <= 0 {
		encounter.ActiveEntityOrder = removeEntityFromOrder(encounter.ActiveEntity.Id, encounter.ActiveEntityOrder)

		for y, entitiesRow := range encounter.Board.Entities.Entities {
			for x, entity := range entitiesRow.Entities {
				if entity.Id == encounter.ActiveEntity.Id {
					encounter.Board.Entities.Entities[y].Entities[x] = &deviant.Entity{}
				}
			}
		}
	}

	for _, entitiesRow := range encounter.Board.Entities.Entities {
		for _, entity := range entitiesRow.Entities {
			if entity.Id == newActiveEntityID {
				encounter.ActiveEntity = entity
			}
		}
	}

	if logger != nil {
		logger.Debug("Updated Active Entity",
			zap.String("actionID", "UpdateActiveEntity"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return true
}
