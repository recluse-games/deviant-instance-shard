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
	} else {
		apCostX = moveAction.FinalXPosition - moveAction.StartXPosition
	}

	if moveAction.StartYPosition > moveAction.FinalYPosition {
		apCostY = moveAction.StartYPosition - moveAction.FinalYPosition
	} else {
		apCostY = moveAction.FinalYPosition - moveAction.StartYPosition
	}

	// Apply all state changes to entity in encounter as well as the activeEntity.
	encounter.Board.Entities.Entities[moveAction.FinalXPosition].Entities[moveAction.FinalYPosition] = encounter.ActiveEntity
	encounter.Board.Entities.Entities[moveAction.StartXPosition].Entities[moveAction.StartYPosition] = &deviant.Entity{}
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - apCostX
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - apCostY

	return true
}
