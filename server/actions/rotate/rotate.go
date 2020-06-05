package rotate

import (
	"fmt"

	"github.com/golang/glog"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

// Rotate Rotates an Entity Based on a Tile Selection
func Rotate(encounter *deviant.Encounter, entityRotateAction *deviant.EntityRotateAction) bool {
	encounter.ActiveEntity.Rotation = entityRotateAction.Rotation

	message := fmt.Sprintf("Action: Rotate: %v", encounter.ActiveEntity.Rotation)
	glog.Info(message)

	return true
}
