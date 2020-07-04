package actions

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestGrantAp(t *testing.T) {
	entityWithNoAp := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Ap:    0,
			MaxAp: 5,
		},
	}

	if GrantAp(entityWithNoAp) != true {
		t.Fail()
	}

	if entityWithNoAp.ActiveEntity.Ap != 5 {
		t.Fail()
	}
}
