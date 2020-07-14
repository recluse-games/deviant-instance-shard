package actions

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// Move Moves the active entity to a desired location and updates the AP to reflect the cost.
func Move(encounter *deviant.Encounter, moveAction *deviant.EntityMoveAction, logger *zap.SugaredLogger) bool {
	var apCostX int32
	var apCostY int32

	if moveAction.StartXPosition > moveAction.FinalXPosition {
		apCostX = moveAction.StartXPosition - moveAction.FinalXPosition
	} else if moveAction.StartXPosition < moveAction.FinalXPosition {
		apCostX = moveAction.FinalXPosition - moveAction.StartXPosition
	} else {
		apCostX = 0
	}

	if moveAction.StartYPosition > moveAction.FinalYPosition {
		apCostY = moveAction.StartYPosition - moveAction.FinalYPosition
	} else if moveAction.StartYPosition < moveAction.FinalYPosition {
		apCostY = moveAction.FinalYPosition - moveAction.StartYPosition
	} else {
		apCostY = 0
	}

	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - apCostX
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - apCostY

	encounter.Board.Entities.Entities[moveAction.StartXPosition].Entities[moveAction.StartYPosition] = &deviant.Entity{}
	encounter.Board.Entities.Entities[moveAction.FinalXPosition].Entities[moveAction.FinalYPosition] = encounter.ActiveEntity

	logger.Debug("Entity Move Processed",
		zap.String("actionID", "Move"),
		zap.String("entityID", encounter.ActiveEntity.Id),
	)

	return true
}
