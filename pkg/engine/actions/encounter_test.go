package actions

import (
	"testing"

	deviant "github.com/recluse-games/deviant-protobuf/genproto/go"
)

func TestProcessWinConditions(t *testing.T) {
	entity := &deviant.Entity{
		Ap:        0,
		MaxAp:     5,
		Alignment: deviant.Alignment_FRIENDLY,
	}

	encounter := &deviant.Encounter{
		Board: &deviant.Board{
			Entities: &deviant.Entities{
				Entities: []*deviant.EntitiesRow{
					{
						Entities: []*deviant.Entity{
							entity,
						},
					},
				},
			},
		},
	}

	isEncounterComplete := ProcessWinConditions(encounter)

	if isEncounterComplete == false {
		t.Fail()
	}

	if encounter.WinningAlignment != deviant.Alignment_FRIENDLY {
		t.Logf("%v", encounter)
		t.Fail()
	}
}
