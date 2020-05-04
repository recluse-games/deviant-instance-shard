package move

import (
	"log"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Move draws a card for entity.
func Move(encounter *deviant.Encounter, moveAction *deviant.EntityMoveAction) bool {
	log.Output(1, "test move update")

	totalApCost := moveAction.FinalXPosition - moveAction.StartXPosition + moveAction.FinalYPosition - moveAction.StartYPosition

	// Apply all state changes to entity in encounter as well as the activeEntity
	encounter.Board.Entities.Entities[moveAction.FinalYPosition].Entities[moveAction.FinalXPosition] = encounter.ActiveEntity
	encounter.Board.Entities.Entities[moveAction.StartXPosition].Entities[moveAction.StartYPosition] = &deviant.Entity{}
	encounter.ActiveEntity.Ap = encounter.ActiveEntity.Ap - totalApCost

	return true
}
