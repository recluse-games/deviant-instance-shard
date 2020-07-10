package rules

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestValidateDraw(t *testing.T) {
	entity := &deviant.Entity{
		Deck: enginetest.GenerateDeckLiteral(5),
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entity,
	}

	if ValidateDraw(encounter, nil) != true {
		t.Logf("Failed to determine a valid draw.")
		t.Fail()
	}

	entity = &deviant.Entity{
		Deck: enginetest.GenerateDeckLiteral(0),
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entity,
	}

	if ValidateDraw(encounter, nil) != false {
		t.Logf("Failed to determine an invalid draw.")
		t.Fail()
	}
}
