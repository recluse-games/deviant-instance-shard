package actions

import "go.uber.org/zap"

//findSliceIndex Locates the index of a particular string in a slice.
func findSliceIndex(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

//removeEntityFromOrder Removes an entityID from the entityTurnOrder of an encounter.
func removeEntityFromOrder(entityID string, slice []string) []string {
	entityIDIndex := findSliceIndex(slice, entityID)

	if len(slice) > entityIDIndex+1 {
		return slice[:entityIDIndex+copy(slice[entityIDIndex:], slice[entityIDIndex+1:])]
	} else if len(slice) == entityIDIndex && len(slice) > 1 {
		return slice[:entityIDIndex+copy(slice[entityIDIndex:], slice[entityIDIndex-1:])]
	}

	if 0 > entityIDIndex-1 && len(slice) >= 1 {
		return slice[:0+copy(slice[0:], slice[1:])]
	}

	return []string{}

}

// GetLogger Returns a zap logger for this package.
func GetLogger() *zap.Logger {
	logger, _ := zap.NewProduction()

	return logger
}
