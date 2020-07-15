package engineutil

import (
	"errors"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
)

// Direction represents a direction
type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

// String returns the string value of the Direction
func (s Direction) String() string {
	strings := [...]string{"UP", "DOWN", "LEFT", "RIGHT"}

	return strings[s]
}

// Location returns the Location value of the Direction
func (s Direction) Location() Location {
	locations := [...]Location{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	return locations[s]
}

// ConvertToDirection Converts a deviant.Direction to a Direction
func ConvertToDirection(deviantDirection deviant.Direction) (*Direction, error) {
	switch deviantDirection {
	case deviant.Direction_UP:
		up := UP

		return &up, nil
	case deviant.Direction_DOWN:
		down := DOWN

		return &down, nil
	case deviant.Direction_LEFT:
		left := LEFT

		return &left, nil
	case deviant.Direction_RIGHT:
		right := RIGHT

		return &right, nil
	default:
		return nil, errors.New("No converted direction exists for this deviant direction")
	}
}
