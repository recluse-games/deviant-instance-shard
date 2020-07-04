package rules

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// ValidateDraw Validates that we can draw from deck.
func ValidateDraw(encounter *deviant.Encounter, logger *zap.Logger) bool {
	if len(encounter.ActiveEntity.Deck.Cards) > 0 {
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
