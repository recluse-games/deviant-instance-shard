package rules

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// ValidateDraw Validates that we can draw from deck.
func ValidateDraw(encounter *deviant.Encounter, logger *zap.SugaredLogger) bool {
	status := false

	if len(encounter.ActiveEntity.Deck.Cards) > 0 {
		status = true
	}

	logger.Debug("Validated Draw",
		"actionID", "ValidateDraw",
		"entityID", encounter.ActiveEntity.Id,
		"succeeded", status,
	)

	return status
}
