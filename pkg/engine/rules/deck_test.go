package rules

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
	"go.uber.org/zap/zaptest"
)

func TestValidateDraw(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entity := &deviant.Entity{
		Id:   "0001",
		Deck: enginetest.GenerateDeckLiteral(5),
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entity,
	}

	if ValidateDraw(encounter, logger.Sugar()) != true {
		t.Logf("Failed to determine a valid draw.")
		t.Fail()
	}

	entity = &deviant.Entity{
		Deck: enginetest.GenerateDeckLiteral(0),
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entity,
	}

	if ValidateDraw(encounter, logger.Sugar()) != false {
		t.Logf("Failed to determine an invalid draw.")
		t.Fail()
	}
}
