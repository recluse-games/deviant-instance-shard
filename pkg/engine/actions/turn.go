package actions

import (
	"github.com/recluse-games/deviant-instance-shard/pkg/engine/engineutil"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap"
)

// GrantAp Grants the currently active entity their maximum AP
func GrantAp(encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.MaxAp

	logger.Debug("Entity Granted AP",
		zap.String("actionID", "GrantAP"),
		zap.String("entityID", encounter.ActiveEntity.Id),
	)

	return true
}

// ChangePhase Updates the current turn phase to the next turn phase.
func ChangePhase(encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
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

	logger.Debug("Phase Changed",
		"actionID", "ChangePhase",
		"entityID", encounter.ActiveEntity.Id,
		"phaseID", encounter.Turn.Phase.String(),
	)

	return true
}

// UpdateActiveEntity Updates the active entity to the next entity in the active entity order.
func UpdateActiveEntity(encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	var newActiveEntityID string

	newActiveEntityIndex, err := engineutil.IndexString(encounter.ActiveEntityOrder, encounter.ActiveEntity.Id)
	if err != nil {
		return false
	}

	if len(encounter.ActiveEntityOrder) > *newActiveEntityIndex+1 {
		newActiveEntityID = encounter.ActiveEntityOrder[*newActiveEntityIndex+1]
	} else if len(encounter.ActiveEntityOrder) == 1 {
		newActiveEntityID = encounter.ActiveEntityOrder[0]
	}

	if encounter.ActiveEntity.Hp <= 0 {
		encounter.ActiveEntityOrder, err = engineutil.RemoveString(encounter.ActiveEntity.Id, encounter.ActiveEntityOrder)
		if err != nil {
			return false
		}

		location, err := engineutil.LocateEntity(encounter.ActiveEntity.Id, encounter.Board)
		if err != nil {
			return false
		}

		encounter.Board.Entities.Entities[location.Y].Entities[location.X] = &deviant.Entity{}
	}

	newActiveEntity, err := engineutil.GetEntity(newActiveEntityID, encounter.Board)
	if err != nil {
		return false
	}

	encounter.ActiveEntity = newActiveEntity

	logger.Debug("Updated Active Entity",
		zap.String("actionID", "UpdateActiveEntity"),
		zap.String("entityID", encounter.ActiveEntity.Id),
	)

	return true
}
