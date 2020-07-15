package actions

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap"
)

// Rotate Rotates an Entity Based on a Tile Selection
func Rotate(encounter *deviant.Encounter, entityRotateAction *deviant.EntityRotateAction, logger *zap.SugaredLogger) bool {
	encounter.ActiveEntity.Rotation = entityRotateAction.Rotation

	logger.Debug("Entity Rotation Processed",
		"actionID", "Rotate",
		"entityID", encounter.ActiveEntity.Id,
		"rotationID", encounter.ActiveEntity.Rotation.String(),
	)

	return true
}
