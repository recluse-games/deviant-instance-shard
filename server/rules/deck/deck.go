package deck

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// ValidateDraw Validates that we can draw from deck.
func ValidateDraw(entity *deviant.Entity) bool {
	if len(entity.Deck.Cards) > 0 {
		return true
	}

	return false
}
