package actions

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestRotate(t *testing.T) {
	encounter := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Id: "0001",
			Ap:    0,
			MaxAp: 5,
			Rotation: deviant.EntityRotationNames_NORTH,
		},
	}

	entityRotateAction := &deviant.EntityRotateAction{
		Id: "0001",
		Rotation: deviant.EntityRotationNames_SOUTH,
	}

	rotate := Rotate(encounter, entityRotateAction, nil)
	if rotate != true {
		t.Logf("Failed to rotate entity")
		t.Fail()
	}
}
