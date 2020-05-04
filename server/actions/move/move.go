package move

import (
	"log"
	"math"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Move draws a card for entity.
func Move(encounter *deviant.Encounter, moveAction *deviant.EntityMoveAction) bool {
	log.Output(1, "test move update")

	// Get the absolute value of the difference between all of our positions on the board.
	totalApCost := math.Abs(float64(moveAction.FinalXPosition - moveAction.StartXPosition + moveAction.FinalYPosition - moveAction.StartYPosition))

	// Apply all state changes to entity in encounter as well as the activeEntity.
	encounter.Board.Entities.Entities[moveAction.FinalYPosition].Entities[moveAction.FinalXPosition] = encounter.ActiveEntity
	encounter.Board.Entities.Entities[moveAction.StartXPosition].Entities[moveAction.StartYPosition] = &deviant.Entity{}
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - int32(totalApCost)

	return true
}
