package rules

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// ValidateSize Checks size of entity hand to ensure we can draw more cards.
func ValidateSize(encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	status := false

	if len(encounter.ActiveEntity.Hand.Cards) <= 6 {
		status = true
	}

	logger.Debug("Validated Draw",
		"actionID", "ValidateDraw",
		"entityID", encounter.ActiveEntity.Id,
		"succeeded", status,
	)

	return status
}
