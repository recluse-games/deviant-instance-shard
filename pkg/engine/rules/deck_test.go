package rules

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestValidateDraw(t *testing.T) {

	entityWithEnoughCardsInDeckToDraw := &deviant.Entity{
		Deck: enginetest.GenerateDeckLiteral(5),
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entityWithEnoughCardsInDeckToDraw,
	}

	if ValidateDraw(encounter, nil) != true {
		t.Fail()
	}

	entityWithoutEnoughCardsInDeckToDraw := &deviant.Entity{
		Deck: enginetest.GenerateDeckLiteral(0),
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entityWithoutEnoughCardsInDeckToDraw,
	}

	if ValidateDraw(encounter, nil) != false {
		t.Fail()
	}
}
