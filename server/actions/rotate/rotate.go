package rotate

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Rotate Rotates an Entity Based on a Tile Selection
func Rotate(encounter *deviant.Encounter, moveAction *deviant.EntityMoveAction) bool {
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - apCostX
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - apCostY

	// Apply all state changes to entity in encounter as well as the activeEntity.
	encounter.Board.Entities.Entities[moveAction.FinalXPosition].Entities[moveAction.FinalYPosition] = encounter.ActiveEntity
	encounter.Board.Entities.Entities[moveAction.StartXPosition].Entities[moveAction.StartYPosition] = &deviant.Entity{}

	return true
}
