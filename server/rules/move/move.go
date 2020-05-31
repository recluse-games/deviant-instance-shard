package move

import (
	"fmt"

	"github.com/golang/glog"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// ValidateApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidateApCost(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction, encounter *deviant.Encounter) bool {
	var totalApCost int32
	var apCostX int32
	var apCostY int32

	if requestedMoveAction.StartXPosition > requestedMoveAction.FinalXPosition {
		apCostX = requestedMoveAction.StartXPosition - requestedMoveAction.FinalXPosition
	} else if requestedMoveAction.StartXPosition < requestedMoveAction.FinalXPosition {
		apCostX = requestedMoveAction.FinalXPosition - requestedMoveAction.StartXPosition
	} else {
		apCostX = 0
	}

	if requestedMoveAction.StartYPosition > requestedMoveAction.FinalYPosition {
		apCostY = requestedMoveAction.StartYPosition - requestedMoveAction.FinalYPosition
	} else if requestedMoveAction.StartYPosition < requestedMoveAction.FinalYPosition {
		apCostY = requestedMoveAction.FinalYPosition - requestedMoveAction.StartYPosition
	} else {
		apCostY = 0
	}

	totalApCost = apCostX + apCostY

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
