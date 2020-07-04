package rules

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// ValidateSize Checks size of entity hand to ensure we can draw more cards.
func ValidateSize(encounter *deviant.Encounter, logger *zap.Logger) bool {
	if len(encounter.ActiveEntity.Hand.Cards) <= 6 {
		return true
	}

	if logger != nil {
		logger.Debug("Validated Draw",
			zap.String("actionID", "ValidateDraw"),
			zap.String("entityID", encounter.ActiveEntity.Id),
		)
	}

	return false
}
