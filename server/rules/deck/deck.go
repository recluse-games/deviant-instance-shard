package deck

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// ValidateDraw Validates that we can draw from deck.
func ValidateDraw(encounter *deviant.Encounter) bool {
	if len(encounter.ActiveEntity.Deck.Cards) > 0 {
		return true
	}

	return false
}
