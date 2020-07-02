package move

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Move draws a card for entity.
func Move(encounter *deviant.Encounter, moveAction *deviant.EntityMoveAction) bool {
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

	// Apply all state changes to entity in encounter as well as the activeEntity.
	encounter.Board.Entities.Entities[moveAction.StartXPosition].Entities[moveAction.StartYPosition] = &deviant.Entity{}
	encounter.Board.Entities.Entities[moveAction.FinalXPosition].Entities[moveAction.FinalYPosition] = encounter.ActiveEntity

	return true
}
