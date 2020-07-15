package engineutil

import (
	"math"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
)

// Location represents a point in 2d space
type Location struct {
	X int32
	Y int32
}

// Add Adds a new location to this location.
func (l *Location) Add(location Location) {
	l.X = l.X + location.X
	l.Y = l.Y + location.Y
}

// Clone Returns a copy of this location.
func (l *Location) Clone() *Location {
	return &Location{X: l.X, Y: l.Y}
}

// Move Moves a location in a direction.
func (l *Location) Move(direction Direction, distance int32) {
	switch direction {
	case UP:
		for i := int32(0); i < distance; i++ {
			l.Add(UP.Location())
		}
	case DOWN:
		for i := int32(0); i < distance; i++ {
			l.Add(DOWN.Location())
		}
	case LEFT:
		for i := int32(0); i < distance; i++ {
			l.Add(LEFT.Location())
		}
	case RIGHT:
		for i := int32(0); i < distance; i++ {
			l.Add(RIGHT.Location())
		}
	}
}

// Rotate a cartesian point around a given cartesian origin point on a 2D plane based on a degree.
func (l *Location) Rotate(origin Location, degree float64) {
	// Convert everything to floating point.
	oxf := float64(origin.X)
	oxy := float64(origin.Y)
	pxf := float64(l.X)
	pyf := float64(l.Y)

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

	l.X = x
	l.Y = y
}

// GetDegree Converts entity rotations into floating point degree representations.
func GetDegree(rotation deviant.EntityRotationNames) float64 {
	switch rotation {
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

// RotateLocations Generated a list of rotated locations around an origin based on a degree
func RotateLocations(origin Location, locations []*Location, rotation deviant.EntityRotationNames) {
	rotationDegree := GetDegree(rotation)

	for _, location := range locations {
		location.Rotate(origin, rotationDegree)
	}
}
