package engineutil

import (
	"errors"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

//IndexString Finds the index of a given string in a slice of strings.
func IndexString(ss []string, s string) (*int, error) {
	for i, n := range ss {
		if s == n {
			return &i, nil
		}
	}
	return nil, errors.New("String entry not found in slice")
}

// LocateEntity Returns the location of an Entity on the board from it's ID.
func LocateEntity(id string, board *deviant.Board) (*Location, error) {
	if board.Entities == nil {
		return nil, errors.New("No entities exist on board")
	}

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

// RemoveString Removes an entityID from the entityTurnOrder of an encounter.
func RemoveString(value string, slice []string) ([]string, error) {
	i, err := IndexString(slice, value)
	if err != nil {
		return slice, errors.New("string does not exist in slice")
	}

	if len(slice) > 1 {
		return append(slice[:*i], slice[*i+1:]...), nil
	}

	return nil, nil
}
