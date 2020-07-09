package actions

import (
	"errors"
	"math"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap"
)

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
func (s Direction) location() location {
	locations := [...]location{location{1, 0}, location{-1, 0}, location{0, 1}, location{0, -1}}

	return locations[s]
}

// Location represents a point in 2d space
type location struct {
	X int32
	Y int32
}

// locateEntity Returns the location of an Entity on the board from it's ID.
func locateEntity(id string, board *deviant.Board) (*location, error) {
	for x, row := range board.Entities.Entities {
		for y, entity := range row.Entities {
			if entity.Id == id {
				return &location{X: int32(x), Y: int32(y)}, nil
			}
		}
	}

	return nil, errors.New("Failed to locate entity with provided ID in Board")
}

// getEntity Returns a pointer to an entity from the board based on it's ID.
func getEntity(id string, board *deviant.Board) (*deviant.Entity, error) {
	for _, row := range board.Entities.Entities {
		for _, entity := range row.Entities {
			if entity.Id == id {
				return entity, nil
			}
		}
	}

	return nil, errors.New("Failed to locate entity with provided ID in Board")
}

// addLocations Adds two locations together and returns the sum
func addLocations(one location, two location) location {
	return location{one.X + two.X, one.Y + two.Y}
}

// offSetStart Generates the starting location for offsets to be calculated from a given action
func moveLocation(start location, direction deviant.Direction, distance int32) location {
	switch direction {
	case deviant.Direction_UP:
		for i := int32(0); i < distance; i++ {
			start = addLocations(start, UP.location())
		}
	case deviant.Direction_DOWN:
		for i := int32(0); i < distance; i++ {
			start = addLocations(start, DOWN.location())
		}
	case deviant.Direction_LEFT:
		for i := int32(0); i < distance; i++ {
			start = addLocations(start, LEFT.location())
		}
	case deviant.Direction_RIGHT:
		for i := int32(0); i < distance; i++ {
			start = addLocations(start, RIGHT.location())
		}
	}

	return start
}

// rotationDegrees Converts entity rotations into floating point degree representations.
func rotationDegrees(characterRotation deviant.EntityRotationNames) float64 {
	switch characterRotation {
	case deviant.EntityRotationNames_NORTH:
		return 180.00
	case deviant.EntityRotationNames_SOUTH:
		return 0.00
	case deviant.EntityRotationNames_EAST:
		return 270.00
	case deviant.EntityRotationNames_WEST:
		return 90.00
	}

	return 0.00
}

// Rotates a cartesian point around a given cartesian origin point on a 2D plane based on a degree.
func rotateLocation(origin location, target location, degree float64) location {
	// Convert everything to floating point.
	oxf := float64(origin.X)
	oxy := float64(origin.Y)
	pxf := float64(target.X)
	pyf := float64(target.Y)

	// Get the radians of our degree and take the sin and cos
	radians := (math.Pi / 180) * degree
	s := math.Sin(radians)
	c := math.Cos(radians)

	// translate point back to origin:
	pxf -= oxf
	pyf -= oxy

	// rotate point
	xnew := pxf*c - pyf*s
	ynew := pxf*s + pyf*c

	// translate point back:
	pxf = xnew + oxf
	pyf = ynew + oxy

	// Convert everything back to ints.
	x := int32(math.RoundToEven(pxf))
	y := int32(math.RoundToEven(pyf))

	return location{x, y}
}

// rotateLocations Generated a list of rotated locations around an origin based on a degree
func rotateLocations(origin location, cardLocations []location, rotation deviant.EntityRotationNames) []location {
	rotatedLocations := []location{}
	rotationDegree := rotationDegrees(rotation)

	for _, location := range cardLocations {
		rotatedLocations = append(rotatedLocations, rotateLocation(origin, location, rotationDegree))
	}

	return rotatedLocations
}

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
