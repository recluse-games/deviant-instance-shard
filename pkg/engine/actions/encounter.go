package actions

import (
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

// ProcessWinConditions Validates if any enemies remain if they don't the encounter will be set to completed.
func ProcessWinConditions(encounter *deviant.Encounter, logger *zap.Logger) bool {
	friendlyCounter := 0
	enemyCounter := 0

	for _, outerValue := range encounter.Board.Entities.Entities {
		for _, innerValue := range outerValue.Entities {
			if innerValue.Alignment == deviant.Alignment_UNFRIENDLY {
				enemyCounter++
			}
		}
	}

	for _, outerValue := range encounter.Board.Entities.Entities {
		for _, innerValue := range outerValue.Entities {
			if innerValue.Alignment == deviant.Alignment_FRIENDLY {
				friendlyCounter++
			}
		}
	}

	if friendlyCounter != 0 {
		encounter.WinningAlignment = deviant.Alignment_FRIENDLY
		encounter.Completed = true
	}

	if enemyCounter != 0 {
		encounter.WinningAlignment = deviant.Alignment_UNFRIENDLY
		encounter.Completed = true
	}

	if logger != nil {
		logger.Debug("Win Conditions Processed",
			zap.String("actionID", "ProcessWinConditions"),
		)
	}

	return true
}