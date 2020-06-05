package rotate

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Rotate Rotates an Entity Based on a Tile Selection
func Rotate(encounter *deviant.Encounter, entityRotateAction *deviant.EntityRotateAction) bool {
	encounter.ActiveEntity.Rotation = entityRotateAction.Rotation

	return true
}
