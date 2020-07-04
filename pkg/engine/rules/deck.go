package rules

import deviant "github.com/recluse-games/deviant-protobuf/genproto/go"

// ValidateDraw Validates that we can draw from deck.
func ValidateDraw(encounter *deviant.Encounter) bool {
	if len(encounter.ActiveEntity.Deck.Cards) > 0 {
		return true
	}

	return false
}
