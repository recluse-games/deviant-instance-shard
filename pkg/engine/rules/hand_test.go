package rules

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestValidateSize(t *testing.T) {

	testEntityWithValidHand := &deviant.Entity{
		Hand: &deviant.Hand{
			Id:    "0000",
			Cards: enginetest.GenerateCardLiterals(1),
		},
	}

	encounter := &deviant.Encounter{
		ActiveEntity: testEntityWithValidHand,
	}

	isHandSizeValid := ValidateSize(encounter)

	if isHandSizeValid != true {
		t.Fail()
	}

	testEntityWithInvalidHand := &deviant.Entity{
		Hand: &deviant.Hand{
			Id:    "0000",
			Cards: enginetest.GenerateCardLiterals(20),
		},
	}

	encounter = &deviant.Encounter{
		ActiveEntity: testEntityWithInvalidHand,
	}

	isHandSizeInvalid := ValidateSize(encounter)

	if isHandSizeInvalid != false {
		t.Fail()
	}
}
