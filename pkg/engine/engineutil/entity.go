package engineutil

import (
	"errors"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

//IndexString Finds the index of a given string in a slice of strings.
func IndexString(a []string, x string) (int, error) {
	for i, n := range a {
		if x == n {
			return i, nil
		}
	}
	return 0, errors.New("String entry not found in slice")
}

// RemoveString Removes a string from a slice based on index.
func RemoveString(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// LocateEntity Returns the location of an Entity on the board from it's ID.
func LocateEntity(id string, board *deviant.Board) (*Location, error) {
	for x, row := range board.Entities.Entities {
		for y, entity := range row.Entities {
			if entity.Id == id {
				return &Location{X: int32(x), Y: int32(y)}, nil
			}
		}
	}

	return nil, errors.New("Failed to locate entity with provided ID in Board")
}

// GetEntity Returns a pointer to an entity from the board based on it's ID.
func GetEntity(id string, board *deviant.Board) (*deviant.Entity, error) {
	for _, row := range board.Entities.Entities {
		for _, entity := range row.Entities {
			if entity.Id == id {
				return entity, nil
			}
		}
	}

	return nil, errors.New("Failed to locate entity with provided ID in Board")
}

// RemoveEntityID Removes an entityID from the entityTurnOrder of an encounter.
func RemoveEntityID(entityID string, slice []string) ([]string, error) {
	entityIndex, err := IndexString(slice, entityID)

	if err != nil {
		return slice, errors.New("ID does not exist in slice")
	}

	if len(slice) > 1 {
		return RemoveString(slice, entityIndex), nil
	}

	return nil, nil
}
