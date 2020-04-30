package turn

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto"
)

func TestValidateEntityAction(t *testing.T) {
	isEntityPlayActionValidInPlayPhase := ValidateEntityAction(deviant.EntityActionNames_PLAY, deviant.TurnPhaseNames_PHASE_ACTION)

	if isEntityPlayActionValidInPlayPhase != true {
		t.Fail()
	}

	isEntityDiscardActionValidInPlayPhase := ValidateEntityAction(deviant.EntityActionNames_DISCARD, deviant.TurnPhaseNames_PHASE_ACTION)

	if isEntityDiscardActionValidInPlayPhase != false {
		t.Fail()
	}

	isDiscardActionValidInDiscardPhase := ValidateEntityAction(deviant.EntityActionNames_DISCARD, deviant.TurnPhaseNames_PHASE_DISCARD)

	if isDiscardActionValidInDiscardPhase != true {
		t.Fail()
	}

	isPlayActionValidInDiscardPhase := ValidateEntityAction(deviant.EntityActionNames_PLAY, deviant.TurnPhaseNames_PHASE_DISCARD)

	if isPlayActionValidInDiscardPhase != false {
		t.Fail()
	}
}
