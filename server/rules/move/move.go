package move

import deviant "github.com/recluse-games/deviant-protobuf/genproto/go"

// ValidateApCost Determines that the entity has the correct amount of AP to perform the requested move.
func ValidateApCost(activeEntity *deviant.Entity, requestedMoveAction *deviant.EntityMoveAction) bool {
	totalApCost := requestedMoveAction.FinalXPosition - requestedMoveAction.StartXPosition + requestedMoveAction.FinalYPosition - requestedMoveAction.StartYPosition
	if totalApCost < activeEntity.Ap {
		return true
	}

	return false
}
