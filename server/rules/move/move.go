package move

import (
	"fmt"

	"github.com/golang/glog"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// ValidateApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidateApCost(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter) bool {
	totalApCost := requestedMoveAction.FinalXPosition - requestedMoveAction.StartXPosition + requestedMoveAction.FinalYPosition - requestedMoveAction.StartYPosition
	if totalApCost <= activeEntity.Ap {
		return true
	}

	return false
}

// ValidateNewLocationEmpty Determines that the entity can actually move to this location on the board.
func ValidateNewLocationEmpty(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter) bool {
	// Validate that the new location is empty
	if encounter.Board.Entities.Entities[requestedMoveAction.FinalXPosition].Entities[requestedMoveAction.FinalYPosition].Id == "" {
		message := fmt.Sprintf("Rule: ValidateNewLocationEmpty: Success")
		glog.Info(message)

		return true
	}

	message := fmt.Sprintf("Rule: ValidateNewLocationEmpty: Failed")
	glog.Error(message)
	return false
}

// ValidateNewLocationSide Determines that the entity can actually move to this location on the board.
func ValidateNewLocationSide(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter) bool {
	if activeEntity.Alignment == deviant.Alignment_FRIENDLY {
		if requestedMoveAction.FinalXPosition <= 3 {
			return true
		}
	}
	if activeEntity.Alignment == deviant.Alignment_UNFRIENDLY {
		if requestedMoveAction.FinalXPosition >= 4 {
			return true
		}
	}

	message := fmt.Sprintf("Rule: ValidateNewLocationSide: Failed")
	glog.Info(message)
	return false
}
