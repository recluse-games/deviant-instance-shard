package rules

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap/zaptest"
)

func TestValidateSize(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Hand: &deviant.Hand{
			Id:    "0000",
			Cards: enginetest.GenerateCardLiterals(1),
		},
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entity,
	}

	isHandSizeValid := ValidateSize(encounter, logger.Sugar())

	if isHandSizeValid != true {
		t.Logf("Failed to detect valid hand size.")
		t.Fail()
	}

	entity = &deviant.Entity{
		Hand: &deviant.Hand{
			Id:    "0000",
			Cards: enginetest.GenerateCardLiterals(20),
		},
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entity,
	}

	isHandSizeInvalid := ValidateSize(encounter, logger.Sugar())

	if isHandSizeInvalid != false {
		t.Logf("Failed to detect invalid hand size.")
		t.Fail()
	}
}
