package actions

import (
	"errors"

	"go.uber.org/zap"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

//findSliceIndex Locates the index of a particular string in a slice.
func findSliceIndex(a []string, x string) (int, error) {
	for i, n := range a {
		if x == n {
			return i, nil
		}
	}
	return len(a), errors.New("entry not found in slice")
}

//removeEntityFromOrder Removes an entityID from the entityTurnOrder of an encounter.
func removeEntityFromOrder(entityID string, slice []string) []string {
	entityIDIndex, err := findSliceIndex(slice, entityID)
	if err != nil {
		return slice
	}

	if len(slice) >= 1 {
		return remove(slice, entityIDIndex)
	}
}

// GetLogger Returns a zap logger for this package.
func GetLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger
}
