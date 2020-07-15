package rules

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap/zaptest"
)

func TestProcess(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entityWithEnoughCardsInDeckToDraw := &deviant.Entity{
		Deck: enginetest.GenerateDeckLiteral(5),
		Hand: enginetest.GenerateHandLiteral(5),
	}

	turn := &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	encounter := &deviant.Encounter{
		ActiveEntity: entityWithEnoughCardsInDeckToDraw,
		Turn:         turn,
	}

	// Test Draw Action in Correct Phase
	if Process(encounter, deviant.EntityActionNames_DRAW, nil, nil, logger.Sugar()) != true {
		t.Fail()
	}

	turn = &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_ACTION,
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entityWithEnoughCardsInDeckToDraw,
		Turn:         turn,
	}

	// Test Draw Action in Incorrect Phase
	if Process(encounter, deviant.EntityActionNames_DRAW, nil, nil, logger.Sugar()) == false {
		t.Logf("%s", "Failed attempting to process")

		t.Fail()
	}

	entityWithoutEnoughCardsInDeckToDraw := &deviant.Entity{
		Deck: enginetest.GenerateDeckLiteral(0),
		Hand: enginetest.GenerateHandLiteral(5),
	}

	turn = &deviant.Turn{
		Phase: deviant.TurnPhaseNames_PHASE_DRAW,
	}

	encounter = &deviant.Encounter{
		ActiveEntity: entityWithoutEnoughCardsInDeckToDraw,
		Turn:         turn,
	}

	if Process(encounter, deviant.EntityActionNames_DRAW, nil, nil, logger.Sugar()) != false {
		t.Logf("%s", "Failed attempting to process")
		t.Fail()
	}

}
