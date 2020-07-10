package engineutil

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestLocation(t *testing.T) {
	up := UP.Location()
	if up.X != 1 || up.Y != 0 {
		t.Logf("Failed to get proper location for up")
		t.Fail()
	}

	down := DOWN.Location()
	if down.X != -1 || down.Y != 0 {
		t.Logf("Failed to get proper location for down")
		t.Fail()
	}

	left := LEFT.Location()
	if left.X != 0 || left.Y != 1 {
		t.Logf("Failed to get proper location for left")
		t.Fail()
	}

	right := RIGHT.Location()
	if right.X != 0 || right.Y != -1 {
		t.Logf("Failed to get proper location for right")
		t.Fail()
	}

}

func TestConvertToDirection(t *testing.T) {
	up, err := ConvertToDirection(deviant.Direction_UP)
	if err != nil {
		t.Logf("Failed to convert deviant up direction to up")
		t.Fail()
	}
	if up.String() != "UP" {
		t.Logf("Failed to convert deviant up direction to up")
		t.Fail()
	}

	down, err := ConvertToDirection(deviant.Direction_DOWN)
	if err != nil {
		t.Logf("Failed to convert deviant down direction to down")
		t.Fail()
	}
	if down.String() != "DOWN" {
		t.Logf("Failed to convert deviant up direction to up")
		t.Fail()
	}

	left, err := ConvertToDirection(deviant.Direction_LEFT)
	if err != nil {
		t.Logf("Failed to convert deviant left direction to left")
		t.Fail()
	}
	if left.String() != "LEFT" {
		t.Logf("Failed to convert deviant up direction to up")
		t.Fail()
	}

	right, err := ConvertToDirection(deviant.Direction_RIGHT)
	if err != nil {
		t.Logf("Failed to convert deviant right direction to right")
		t.Fail()
	}
	if right.String() != "RIGHT" {
		t.Logf("Failed to convert deviant up direction to up")
		t.Fail()
	}
}
