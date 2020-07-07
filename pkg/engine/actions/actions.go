package actions

import (
	"errors"

	"go.uber.org/zap"
)

//remove Removes a string from a slice based on index.
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

//findIndex Finds the index of a given string in a slice of strings.
func findIndex(a []string, x string) (int, error) {
	for i, n := range a {
		if x == n {
			return i, nil
		}
	}
	return 0, errors.New("entry not found in slice")
}

//removeEntityID Removes an entityID from the entityTurnOrder of an encounter.
func removeEntityID(entityID string, slice []string) ([]string, error) {
	entityIndex, err := findIndex(slice, entityID)
	if err != nil {
		return slice, errors.New("ID does not exist in slice")
	}

	if len(slice) > 1 {
		return remove(slice, entityIndex), nil
	}

	return nil, nil
}

// GetLogger Returns a zap logger for this package.
func GetLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger
}
