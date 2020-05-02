package encounter

import deviant "github.com/recluse-games/deviant-protobuf/genproto"

// ProcessWinConditions Validates if any enemies remain if they don't the encounter will be set to completed.
func ProcessWinConditions(encounter *deviant.Encounter) bool {
	enemyCounter := 0

	// Apply all state changes to entity in encounter as well as the activeEntity
	for _, outerValue := range encounter.Board.Entities.Entities {
		for _, innerValue := range outerValue.Entities {
			if innerValue.Alignment == deviant.Alignment_UNFRIENDLY {
				enemyCounter++
			}
		}
	}

	if enemyCounter == 0 {
		encounter.Completed = true
	}

	return true
}
