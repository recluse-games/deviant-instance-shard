package hand

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// ValidateSize Checks size of entity hand to ensure we can draw more cards.
func ValidateSize(encounter *deviant.Encounter) bool {
	if len(encounter.ActiveEntity.Hand.Cards) <= 6 {
		return true
	}

	return false
}
