package actions

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// Rotate Rotates an Entity Based on a Tile Selection
func Rotate(encounter *deviant.Encounter, entityRotateAction *deviant.EntityRotateAction, logger *zap.Logger) bool {
	encounter.ActiveEntity.Rotation = entityRotateAction.Rotation

	if logger != nil {
		logger.Debug("Entity Rotation Processed",
			zap.String("actionID", "Rotate"),
			zap.String("entityID", encounter.ActiveEntity.Id),
			zap.String("rotationID", encounter.ActiveEntity.Rotation.String()),
		)
	}

	return true
}
